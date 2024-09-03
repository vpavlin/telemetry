CREATE TABLE IF NOT EXISTS regularHistoryQuery (
    id SERIAL PRIMARY KEY,
    createdAt INTEGER NOT NULL,
    peerId VARCHAR(255) NOT NULL,
    nodeName VARCHAR(255) NOT NULL,
    timestamp INTEGER NOT NULL,
    statusVersion VARCHAR(31),
    count INTEGER NOT NULL,
    pubsubTopic VARCHAR(255) NOT NULL,
    deviceType VARCHAR(255),
    CONSTRAINT regularHistoryQuery_unique unique(timestamp, peerId, pubsubTopic, count)
);
