package imgexpirer

import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/filanov/bm-inventory/internal/events"
	logutil "github.com/filanov/bm-inventory/pkg/log"
	"github.com/filanov/bm-inventory/pkg/requestid"
	"github.com/sirupsen/logrus"
)

const imagePrefix = "discovery-image-"
const imagePrefixLen = len(imagePrefix)

type Manager struct {
	log           logrus.FieldLogger
	s3Client      *s3.S3
	s3Bucket      string
	deleteTime    time.Duration
	eventsHandler events.Handler
}

func NewManager(log logrus.FieldLogger, s3Client *s3.S3, s3Bucket string, deleteTime time.Duration, eventsHandler events.Handler) *Manager {
	return &Manager{
		log:           log,
		s3Client:      s3Client,
		s3Bucket:      s3Bucket,
		deleteTime:    deleteTime,
		eventsHandler: eventsHandler,
	}
}

func (m *Manager) ExpirationTask() {
	ctx := requestid.ToContext(context.Background(), requestid.NewID())
	log := logutil.FromContext(ctx, m.log)
	now := time.Now()
	prefix := imagePrefix

	log.Info("Image expiration monitor woke up, checking for expired images...")
	// Currently limited to 1000 objects
	output, err := m.s3Client.ListObjects(&s3.ListObjectsInput{Bucket: &m.s3Bucket, Prefix: &prefix})
	if err != nil {
		log.WithError(err).Error("Error listing objects")
		return
	}
	for _, object := range output.Contents {
		m.handleObject(ctx, log, object, now)
	}
}

func (m *Manager) handleObject(ctx context.Context, log logrus.FieldLogger, object *s3.Object, now time.Time) {
	// The timestamp that we really want is stored in a tag, but we check this one first as a cost optimization
	if now.Before(object.LastModified.Add(m.deleteTime)) {
		return
	}
	objectTags, err := m.s3Client.GetObjectTagging(&s3.GetObjectTaggingInput{Bucket: &m.s3Bucket, Key: object.Key})
	if err != nil {
		log.WithError(err).Errorf("Error getting tags for object %s", *object.Key)
		return
	}
	for _, tag := range objectTags.TagSet {
		if *tag.Key == "create_sec_since_epoch" {
			objTime, _ := strconv.ParseInt(*tag.Value, 10, 64)
			if now.Before(time.Unix(objTime, 0).Add(m.deleteTime)) {
				m.deleteObject(ctx, log, object)
			}
		}
	}
}

func (m *Manager) deleteObject(ctx context.Context, log logrus.FieldLogger, object *s3.Object) {
	_, err := m.s3Client.DeleteObject(&s3.DeleteObjectInput{Bucket: &m.s3Bucket, Key: object.Key})
	if err != nil {
		log.WithError(err).Errorf("Error deleting object %s", *object.Key)
		return
	}
	eventMsg := "Deleted image from backend because it expired. It may be generated again at any time."
	m.eventsHandler.AddEvent(ctx, clusterIDFromImageName(*object.Key), eventMsg, time.Now())
	log.Infof("Deleted expired image %s", *object.Key)
}

func clusterIDFromImageName(imgName string) string {
	//Image name format is "discovery-image-<clusterID>"
	return imgName[imagePrefixLen:]
}
