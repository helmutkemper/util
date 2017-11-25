package util

import (
  "fmt"
  "github.com/helmutkemper/gOsm/consts"
  "github.com/helmutkemper/db"
)

func ExampleAutoIdLog_Get_With_Database() {
  db.Connect( consts.DB_SERVER_TEST, consts.DB_DATABASE_TEST + "ExampleAutoIdLog_Get" )

  autoIdCollection    = "gOsm_test_collection_autoIdTest"
  autoIdLogCollection = "gOsm_test_collection_autoIdTestLog"

  AutoId.Prepare( false )

  fmt.Printf( "id: %v\n", AutoId.Get( "onlyATest" ) )
  fmt.Printf( "id: %v\n", AutoId.Get( "onlyATest" ) )
  fmt.Printf( "id: %v\n", AutoId.Get( "onlyATest" ) )

  db.DropDatabase( consts.DB_DATABASE_TEST + "ExampleAutoIdLog_Get" )

  autoIdCollection    = consts.DB_OSM_FILE_PSEUDO_ID_COLLECTIONS
  autoIdLogCollection = consts.DB_OSM_FILE_PSEUDO_ID_COLLECTIONS_LOG

  // Output:
  // id: 9223372036854775805
  // id: 9223372036854775804
  // id: 9223372036854775803
}

func ExampleAutoIdLog_Get_Without_Database() {
  AutoId.Prepare( true )

  fmt.Printf( "id: %v\n", AutoId.Get( "onlyATest" ) )
  fmt.Printf( "id: %v\n", AutoId.Get( "onlyATest" ) )
  fmt.Printf( "id: %v\n", AutoId.Get( "onlyATest" ) )

  // Output:
  // id: 9223372036854775806
  // id: 9223372036854775805
  // id: 9223372036854775804
}
