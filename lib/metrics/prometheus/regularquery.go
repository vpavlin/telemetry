package prometheus

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/status-im/telemetry/lib/common"
	"github.com/status-im/telemetry/pkg/types"
)

func ProcessReqularMissingMsgQuery(db *sql.DB, errs *common.MetricErrors, data *types.PrometheusMetricData) error {
	fmt.Println("Processing regular query")
	for _, m := range data.Value {
		labels := make(map[string]string)
		for _, l := range m.GetLabel() {
			labels[*l.Name] = *l.Value
		}

		fmt.Println(labels)

		if labels["type"] == "relay" {
			stmt, err := db.Prepare("INSERT INTO regularHistoryQuery (timestamp, nodeName, peerId, statusVersion, createdAt, deviceType, count, pubsubTopic) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;")
			if err != nil {
				return err
			}

			defer stmt.Close()

			data.CreatedAt = time.Now().Unix()
			lastInsertId := 0
			err = stmt.QueryRow(
				data.Timestamp,
				data.NodeName,
				data.PeerID,
				data.StatusVersion,
				data.CreatedAt,
				data.DeviceType,
				m.Counter.Value,
				labels["pubsubTopic"],
			).Scan(&lastInsertId)
			if err != nil {
				errs.Append(data.ID, fmt.Sprintf("Error saving regular missing message query metric: %v", err))
				return err
			}
			data.ID = lastInsertId

			return nil
		}
	}

	return nil

}
