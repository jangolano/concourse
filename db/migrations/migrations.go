package migrations

import "github.com/BurntSushi/migration"

var Migrations = []migration.Migrator{
	InitialSchema,
	MoveSourceAndMetadataToVersionedResources,
	AddTypeToVersionedResources,
	RemoveTransitionalCurrentVersions,
	NonNullableVersionInfo,
	AddOneOffNameSequence,
	AddHijackURLToBuilds,
	AddTimestampsToBuilds,
	CreateLocks,
	AddBuildEvents,
	ReplaceBuildsAbortHijackURLsWithGuidAndEndpoint,
	ReplaceBuildEventsIDWithEventID,
	AddLocks,
	DropOldLocks,
	AddConfig,
	AddNameToBuildInputs,
	AddEngineAndEngineMetadataToBuilds,
	AddVersionToBuildEvents,
	AddCompletedToBuilds,
	AddWorkers,
	AddEnabledToBuilds,
	CreateEventIDSequencesForInFlightBuilds,
	AddResourceTypesToWorkers,
	AddPlatformAndTagsToWorkers,
	AddIdToConfig,
	ConvertJobBuildConfigToJobPlans,
	AddCheckErrorToResources,
	AddPausedToResources,
	AddPausedToJobs,
	CreateJobsSerialGroups,
	CreatePipes,
	RenameConfigToPipelines,
	RenamePipelineIDToVersionAddPrimaryKey,
	AddNameToPipelines,
	AddPipelineIDToResources,
	AddPipelineIDToJobs,
	AddPausedToPipelines,
	AddOrderingToPipelines,
	AddInputsDeterminedToBuilds,
	AddExplicitToBuildOutputs,
}
