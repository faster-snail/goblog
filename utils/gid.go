package utils

import (
	"fmt"
	"sync"
	"time"
	"log"
)

type iDWorker struct {
	startTime             int64
	workerIDBits          uint
	datacenterIDBits      uint
	maxWorkerID           int64
	maxDatacenterID       int64
	sequenceBits          uint
	workerIDLeftShift     uint
	datacenterIDLeftShift uint
	timestampLeftShift    uint
	sequenceMask          int64
	workerID              int64
	datacenterID          int64
	sequence              int64
	lastTimestamp         int64
	signMask              int64
	idLock                *sync.Mutex
}

//NewID id生成
func NewID () int64 {
	var id iDWorker
	err := id.initIDWorker(96,59)
	if err != nil {
		log.Fatal(err)
	}
	newID,_ := id.nextID()
	return newID
}


func (iw *iDWorker) initIDWorker(workerID, datacenterID int64) error {

	var baseValue int64 = -1
	iw.startTime = 1563691276693
	iw.workerIDBits = 5
	iw.datacenterIDBits = 5
	iw.maxWorkerID = baseValue ^ (baseValue << iw.workerIDBits)
	iw.maxDatacenterID = baseValue ^ (baseValue << iw.datacenterIDBits)
	iw.sequenceBits = 12
	iw.workerIDLeftShift = iw.sequenceBits
	iw.datacenterIDLeftShift = iw.workerIDBits + iw.workerIDLeftShift
	iw.timestampLeftShift = iw.datacenterIDBits + iw.datacenterIDLeftShift
	iw.sequenceMask = baseValue ^ (baseValue << iw.sequenceBits)
	iw.sequence = 0
	iw.lastTimestamp = -1
	iw.signMask = ^baseValue + 1

	iw.idLock = &sync.Mutex{}

	if iw.workerID < 0 || iw.workerID > iw.maxWorkerID {
		return fmt.Errorf("workerID[%v] "+
		"is less than 0 or greater than maxWorkerID[%v].", workerID, datacenterID)
	}
	if iw.datacenterID < 0 || iw.datacenterID > iw.maxDatacenterID {
		return fmt.Errorf("datacenterID[%d] "+
		"is less than 0 or greater than maxDatacenterID[%d].", workerID, datacenterID)
	}
	iw.workerID = workerID
	iw.datacenterID = datacenterID
	return nil
}

func (iw *iDWorker) nextID() (int64, error) {
	iw.idLock.Lock()
	timestamp := time.Now().UnixNano()
	if timestamp < iw.lastTimestamp {
		return -1, fmt.Errorf("Clock moved backwards.  Refusing to generate id for %d milliseconds", iw.lastTimestamp-timestamp)
	}

	if timestamp == iw.lastTimestamp {
		iw.sequence = (iw.sequence + 1) & iw.sequenceMask
		if iw.sequence == 0 {
			timestamp = iw.tilNextMillis()
			iw.sequence = 0
		}
	} else {
		iw.sequence = 0
	}

	iw.lastTimestamp = timestamp

	iw.idLock.Unlock()

	id := ((timestamp - iw.startTime) << iw.timestampLeftShift) |
		(iw.datacenterID << iw.datacenterIDLeftShift) |
		(iw.workerID << iw.workerIDLeftShift) |
		iw.sequence

	if id < 0 {
		id = -id
	}

	return id, nil
}

func (iw *iDWorker) tilNextMillis() int64 {
	timestamp := time.Now().UnixNano()
	if timestamp <= iw.lastTimestamp {
		timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	}
	return timestamp
}