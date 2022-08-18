// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: query.sql

package eventscheduler

import (
	"context"
)

const getTopicMessageIds = `-- name: GetTopicMessageIds :many




SELECT topic, ledgerid, entryid, batchidx, partitionidx FROM pulsar WHERE topic = $1
`

// -- name: UpsertRecord :exec
// INSERT INTO records (id, value, payload) VALUES ($1, $2, $3)
// ON CONFLICT (id) DO UPDATE SET value = EXCLUDED.value, payload = EXCLUDED.payload;
// -- name: UpsertRecords :exec
// INSERT INTO records (id, value, payload)
// SELECT unnest(@ids) AS id,
//
//	unnest(@values) AS names,
//	unnest(@payloads) AS payloads
//
// ON CONFLICT (id) DO UPDATE SET value = EXCLUDED.value, payload = EXCLUDED.payload;
// -- name: UpdateRecord :exec
// UPDATE records SET value = $2, payload = $3 WHERE id = $1;
// -- name: DeleteRecord :exec
// DELETE FROM records WHERE id = $1;
func (q *Queries) GetTopicMessageIds(ctx context.Context, topic string) ([]Pulsar, error) {
	rows, err := q.db.Query(ctx, getTopicMessageIds, topic)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pulsar
	for rows.Next() {
		var i Pulsar
		if err := rows.Scan(
			&i.Topic,
			&i.Ledgerid,
			&i.Entryid,
			&i.Batchidx,
			&i.Partitionidx,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRuns = `-- name: ListRuns :many

SELECT run_id, job_id, executor, assignment, deleted, last_modified FROM runs ORDER BY run_id
`

// -- name: GetRecord :one
// SELECT * FROM records WHERE id = $1 LIMIT 1;
func (q *Queries) ListRuns(ctx context.Context) ([]Run, error) {
	rows, err := q.db.Query(ctx, listRuns)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Run
	for rows.Next() {
		var i Run
		if err := rows.Scan(
			&i.RunID,
			&i.JobID,
			&i.Executor,
			&i.Assignment,
			&i.Deleted,
			&i.LastModified,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const upsertMessageId = `-- name: UpsertMessageId :exec
INSERT INTO pulsar (topic, ledgerId, entryId, batchIdx, partitionIdx) VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (topic) DO UPDATE SET ledgerId = EXCLUDED.ledgerId, entryId = EXCLUDED.entryId, batchIdx = EXCLUDED.batchIdx, partitionIdx = EXCLUDED.partitionIdx
`

type UpsertMessageIdParams struct {
	Topic        string `db:"topic"`
	Ledgerid     int64  `db:"ledgerid"`
	Entryid      int64  `db:"entryid"`
	Batchidx     int32  `db:"batchidx"`
	Partitionidx int32  `db:"partitionidx"`
}

func (q *Queries) UpsertMessageId(ctx context.Context, arg UpsertMessageIdParams) error {
	_, err := q.db.Exec(ctx, upsertMessageId,
		arg.Topic,
		arg.Ledgerid,
		arg.Entryid,
		arg.Batchidx,
		arg.Partitionidx,
	)
	return err
}
