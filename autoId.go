package util

// https://godoc.org/labix.org/v2/mgo#Query.Apply

import (
  log "github.com/helmutkemper/seelog"
  "github.com/helmutkemper/db"
  "github.com/helmutkemper/gOsm/consts"
  "math"
)

var autoIdCollection    string
var autoIdLogCollection string

func init() {
  autoIdCollection    = consts.DB_OSM_FILE_PSEUDO_ID_COLLECTIONS
  autoIdLogCollection = consts.DB_OSM_FILE_PSEUDO_ID_COLLECTIONS_LOG
  idFromTest          = 0
}

var idFromTest int64
var AutoId autoIdLog = autoIdLog{}

type autoIdLog struct{
  Id          int64       `bson:"_id,omitempty"`
  Collection  string      `bson:"collection"`
  *db.DbStt               `bson:"-"`
}

func( el *autoIdLog ) Prepare( dataBaseOffLineLBoo bool ) {
  if dataBaseOffLineLBoo == true {
    idFromTest = math.MaxInt64
    return
  }

  err := el.DbStt.TestConnection()
  if err != nil {
    return
  }

  err = el.DbStt.AutoIncrementPrepare( autoIdCollection )
  if err != nil {
    log.Critical( "gOsm.geoMath.utilAutoId.error: %v", err )
  }
}

func( el *autoIdLog ) Get( collectionNameAStr string ) int64 {
  if idFromTest != 0 {
    idFromTest -= 1
    return idFromTest
  }

  err, id := el.DbStt.AutoIncrement( autoIdCollection )
  if err != nil {
    log.Critical( "gOsm.geoMath.utilAutoId.error: %v", err )
    return 0
  }

  el.Id          = id
  el.Collection  = collectionNameAStr
  err = el.insert( autoIdLogCollection )
  if err != nil {
    log.Critical( "gOsm.geoMath.utilAutoId.error: %v", err )
    return 0
  }

  return id
}

func( el *autoIdLog ) insert( collectionNameAStr string ) error {
  return el.DbStt.Insert( collectionNameAStr, el )
}

