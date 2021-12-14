# aws-meta

## Usage

```bash
aws-meta -h
Usage of aws-meta:
  -r string
    Select a AWS Region (default "us-west-1")
```

```json
Cluster Details:
----------------
{
  DBClusters: [
    {
      ActivityStreamStatus: "stopped",
      AllocatedStorage: 1,
      AutoMinorVersionUpgrade: false,
      AvailabilityZones: ["us-west-1b","us-west-1c","us-west-1a"],
      BackupRetentionPeriod: 35,
      ClusterCreateTime: 2019-08-28 06:23:18.44 +0000 UTC,
      CopyTagsToSnapshot: true,
      CrossAccountClone: false,
      DBClusterArn: "arn:aws:rds:us-west-1:xxxx:cluster:aws-meta-test",
      DBClusterIdentifier: "aws-meta-test",
      DBClusterMembers: [{
          DBClusterParameterGroupStatus: "in-sync",
          DBInstanceIdentifier: "aws-meta-test-1",
          IsClusterWriter: false,
          PromotionTier: 0
        },{
          DBClusterParameterGroupStatus: "in-sync",
          DBInstanceIdentifier: "aws-meta-test-2",
          IsClusterWriter: true,
          PromotionTier: 0
        }],
      DBClusterParameterGroup: "default.aurora-mysql5.7",
      DBSubnetGroup: "default",
      DatabaseName: "aws-meta-test",
      DbClusterResourceId: "cluster-xxxx",
      DeletionProtection: true,
      EarliestRestorableTime: 2021-11-05 11:18:11.154 +0000 UTC,
      EnabledCloudwatchLogsExports: ["slowquery"],
      Endpoint: "aws-test.us-west-1.rds.amazonaws.com",
      Engine: "aurora-mysql",
      EngineMode: "provisioned",
      EngineVersion: "5.7.mysql_aurora.2.07.1",
      HostedZoneId: "default",
      HttpEndpointEnabled: false,
      IAMDatabaseAuthenticationEnabled: false,
      LatestRestorableTime: 2021-12-10 23:23:50.128 +0000 UTC,
      MasterUsername: "aws-meta-test",
      MultiAZ: true,
      Port: 3306,
      PreferredBackupWindow: "11:11-11:41",
      PreferredMaintenanceWindow: "sun:07:46-sun:08:16",
      ReadReplicaIdentifiers: ["arn:aws:rds:us-west-2:xxxxx:cluster:aws-meta-test-replica"],
      ReaderEndpoint: "aws-test-ro.us-west-1.rds.amazonaws.com",
      Status: "available",
      StorageEncrypted: false,
      TagList: [
        {
          Key: "service_portal",
          Value: "https://aws.amazon.com/rds/pricing/pricing-on-demand-mysql-aurora-mysql-on-demand.html"
        },
        {
          Key: "environment",
          Value: "staging"
        },
        {
          Key: "backup_1_year",
          Value: "yes"
        },
        {
          Key: "jira",
          Value: "https://jira.example.com"
        }
      ],
      VpcSecurityGroups: [{
          Status: "active",
          VpcSecurityGroupId: "sg-xxxx"
        }]
    },
```
