package packet

// Code generated from Base/3.8.0src-recommend/src/main/java/emu/lunarcore/server/packet/CmdId.java; DO NOT EDIT.

const (
	NONE                                       uint16 = 0
	SetPlayerInfoCsReq                         uint16 = 3
	UpdatePsnSettingsInfoCsReq                 uint16 = 4
	GetGameStateServiceConfigScRsp             uint16 = 5
	SetGenderCsReq                             uint16 = 6
	GetBasicInfoScRsp                          uint16 = 7
	PlayerLoginScRsp                           uint16 = 9
	AntiAddictScNotify                         uint16 = 10
	PlayerGetTokenScRsp                        uint16 = 11
	SetLanguageScRsp                           uint16 = 15
	GetSecretKeyInfoCsReq                      uint16 = 18
	GmTalkScNotify                             uint16 = 20
	UpdatePlayerSettingCsReq                   uint16 = 22
	ClientObjDownloadDataScNotify              uint16 = 23
	UpdatePsnSettingsInfoScRsp                 uint16 = 24
	GetBasicInfoCsReq                          uint16 = 25
	GateServerScNotify                         uint16 = 27
	FeatureSwitchClosedScNotify                uint16 = 31
	SetGameplayBirthdayCsReq                   uint16 = 32
	QueryProductInfoCsReq                      uint16 = 33
	AceAntiCheaterCsReq                        uint16 = 34
	GetLevelRewardTakenListCsReq               uint16 = 35
	ClientDownloadDataScNotify                 uint16 = 36
	RegionStopScNotify                         uint16 = 37
	PlayerLogoutCsReq                          uint16 = 38
	PlayerHeartBeatCsReq                       uint16 = 39
	ReserveStaminaExchangeCsReq                uint16 = 40
	PlayerForceSyncGameStateFinishCsReq        uint16 = 43
	ClientObjUploadScRsp                       uint16 = 44
	GetLevelRewardCsReq                        uint16 = 46
	PlayerLoginCsReq                           uint16 = 47
	GetVideoVersionKeyScRsp                    uint16 = 48
	ExchangeStaminaScRsp                       uint16 = 50
	PlayerGetTokenCsReq                        uint16 = 51
	ExchangeStaminaCsReq                       uint16 = 52
	PlayerLoginFinishScRsp                     uint16 = 53
	SetPlayerInfoScRsp                         uint16 = 55
	SetLanguageCsReq                           uint16 = 57
	ReserveStaminaExchangeScRsp                uint16 = 59
	GmTalkScRsp                                uint16 = 61
	SetNicknameScRsp                           uint16 = 62
	RetcodeNotify                              uint16 = 65
	QueryProductInfoScRsp                      uint16 = 66
	SetNicknameCsReq                           uint16 = 67
	AceAntiCheaterScRsp                        uint16 = 68
	MonthCardRewardNotify                      uint16 = 72
	ServerAnnounceNotify                       uint16 = 73
	UseReserveAndFuelExchangeStaminaScRsp      uint16 = 74
	StaminaInfoScNotify                        uint16 = 78
	PlayerHeartBeatScRsp                       uint16 = 79
	GetSecretKeyInfoScRsp                      uint16 = 80
	GetAuthkeyCsReq                            uint16 = 81
	PlayerLoginFinishCsReq                     uint16 = 83
	PlayerForceSyncGameStateFinishScRsp        uint16 = 84
	UpdateFeatureSwitchScNotify                uint16 = 85
	SetRedPointStatusScNotify                  uint16 = 86
	SetGenderScRsp                             uint16 = 88
	ClientObjUploadCsReq                       uint16 = 90
	GmTalkCsReq                                uint16 = 91
	GetLevelRewardScRsp                        uint16 = 92
	GetAuthkeyScRsp                            uint16 = 93
	PlayerKickOutScNotify                      uint16 = 95
	GetVideoVersionKeyCsReq                    uint16 = 96
	DailyRefreshNotify                         uint16 = 97
	GetLevelRewardTakenListScRsp               uint16 = 98
	UpdatePlayerSettingScRsp                   uint16 = 99
	SetGameplayBirthdayScRsp                   uint16 = 100
	PVEBattleResultScRsp                       uint16 = 109
	GetCurBattleInfoScRsp                      uint16 = 111
	QuitBattleScRsp                            uint16 = 117
	SyncClientResVersionScRsp                  uint16 = 120
	QuitBattleCsReq                            uint16 = 138
	ServerSimulateBattleFinishScNotify         uint16 = 145
	PVEBattleResultCsReq                       uint16 = 147
	GetCurBattleInfoCsReq                      uint16 = 151
	RebattleByClientCsNotify                   uint16 = 152
	BattleLogReportScRsp                       uint16 = 161
	ReBattleAfterBattleLoseCsNotify            uint16 = 176
	SyncClientResVersionCsReq                  uint16 = 189
	QuitBattleScNotify                         uint16 = 195
	GetPreAvatarGrowthInfoCsReq                uint16 = 301
	GetPreAvatarActivityListScRsp              uint16 = 303
	SetAvatarEnhancedIdScRsp                   uint16 = 306
	SetMultipleAvatarPathsCsReq                uint16 = 307
	GetAvatarDataScRsp                         uint16 = 309
	TakeOffRelicScRsp                          uint16 = 310
	UnlockSkilltreeScRsp                       uint16 = 311
	GrowthTargetAvatarChangedScNotify          uint16 = 313
	UnlockAvatarSkinScNotify                   uint16 = 314
	AddMultiPathAvatarScNotify                 uint16 = 315
	AvatarExpUpScRsp                           uint16 = 317
	PromoteAvatarScRsp                         uint16 = 320
	UnlockAvatarPathScRsp                      uint16 = 325
	MarkAvatarCsReq                            uint16 = 326
	SetPlayerOutfitScRsp                       uint16 = 330
	AvatarSpecialSkilltreeUnlockScNotify       uint16 = 332
	SetAvatarPathCsReq                         uint16 = 333
	DressAvatarSkinCsReq                       uint16 = 335
	AvatarPathChangedNotify                    uint16 = 336
	TakeOffRelicCsReq                          uint16 = 337
	AvatarExpUpCsReq                           uint16 = 338
	TakeOffEquipmentScRsp                      uint16 = 345
	TakeOffAvatarSkinCsReq                     uint16 = 346
	GetAvatarDataCsReq                         uint16 = 347
	GetPreAvatarGrowthInfoScRsp                uint16 = 349
	RankUpAvatarScRsp                          uint16 = 350
	UnlockSkilltreeCsReq                       uint16 = 351
	RankUpAvatarCsReq                          uint16 = 352
	SetPlayerOutfitCsReq                       uint16 = 355
	MarkAvatarScRsp                            uint16 = 357
	TakeOffEquipmentCsReq                      uint16 = 361
	TakePromotionRewardScRsp                   uint16 = 362
	SetAvatarPathScRsp                         uint16 = 366
	TakePromotionRewardCsReq                   uint16 = 367
	SetAvatarEnhancedIdCsReq                   uint16 = 369
	SetGrowthTargetAvatarScRsp                 uint16 = 370
	SetGrowthTargetAvatarCsReq                 uint16 = 373
	AddAvatarScNotify                          uint16 = 376
	DressRelicAvatarCsReq                      uint16 = 381
	UnlockAvatarPathCsReq                      uint16 = 385
	GetPreAvatarActivityListCsReq              uint16 = 388
	PromoteAvatarCsReq                         uint16 = 389
	DressAvatarScRsp                           uint16 = 391
	TakeOffAvatarSkinScRsp                     uint16 = 392
	DressRelicAvatarScRsp                      uint16 = 393
	DressAvatarCsReq                           uint16 = 395
	SetMultipleAvatarPathsScRsp                uint16 = 397
	DressAvatarSkinScRsp                       uint16 = 398
	GetWaypointScRsp                           uint16 = 409
	GetChapterScRsp                            uint16 = 411
	SetCurWaypointScRsp                        uint16 = 417
	BNJCKFMEDCC                                uint16 = 489
	NFOMIJJIGFA                                uint16 = 495
	GetMarkItemListCsReq                       uint16 = 501
	RelicReforgeConfirmCsReq                   uint16 = 502
	SetTurnFoodSwitchScRsp                     uint16 = 503
	SyncTurnFoodNotify                         uint16 = 506
	GetRelicFilterPlanCsReq                    uint16 = 507
	GetBagScRsp                                uint16 = 509
	SellItemCsReq                              uint16 = 510
	LockEquipmentScRsp                         uint16 = 511
	DestroyItemScRsp                           uint16 = 513
	AddEquipmentScNotify                       uint16 = 514
	ComposeLimitNumCompleteNotify              uint16 = 515
	PromoteEquipmentScRsp                      uint16 = 517
	RelicReforgeConfirmScRsp                   uint16 = 518
	MarkItemScRsp                              uint16 = 519
	UseItemScRsp                               uint16 = 520
	GetRecyleTimeCsReq                         uint16 = 526
	RelicFilterPlanClearNameScNotify           uint16 = 527
	DiscardRelicCsReq                          uint16 = 530
	RelicReforgeCsReq                          uint16 = 531
	AddRelicFilterPlanCsReq                    uint16 = 532
	DiscardRelicScRsp                          uint16 = 533
	ModifyRelicFilterPlanCsReq                 uint16 = 534
	ExchangeHcoinCsReq                         uint16 = 535
	LockRelicScRsp                             uint16 = 537
	PromoteEquipmentCsReq                      uint16 = 538
	MarkRelicFilterPlanCsReq                   uint16 = 539
	RelicReforgeScRsp                          uint16 = 541
	ExpUpEquipmentScRsp                        uint16 = 545
	ComposeSelectedRelicCsReq                  uint16 = 546
	GetBagCsReq                                uint16 = 547
	GetMarkItemListScRsp                       uint16 = 549
	ExpUpRelicCsReq                            uint16 = 550
	LockEquipmentCsReq                         uint16 = 551
	ComposeItemScRsp                           uint16 = 552
	GeneralVirtualItemDataNotify               uint16 = 555
	GetRecyleTimeScRsp                         uint16 = 557
	ExpUpEquipmentCsReq                        uint16 = 561
	RechargeSuccNotify                         uint16 = 562
	DeleteRelicFilterPlanScRsp                 uint16 = 565
	SellItemScRsp                              uint16 = 567
	ModifyRelicFilterPlanScRsp                 uint16 = 568
	CancelMarkItemNotify                       uint16 = 569
	DestroyItemCsReq                           uint16 = 570
	DeleteRelicFilterPlanCsReq                 uint16 = 572
	ComposeLimitNumUpdateNotify                uint16 = 573
	ComposeItemCsReq                           uint16 = 576
	MarkItemCsReq                              uint16 = 577
	MarkRelicFilterPlanScRsp                   uint16 = 579
	ExpUpRelicScRsp                            uint16 = 581
	BatchRankUpEquipmentScRsp                  uint16 = 583
	SetTurnFoodSwitchCsReq                     uint16 = 588
	UseItemCsReq                               uint16 = 589
	RankUpEquipmentScRsp                       uint16 = 591
	ComposeSelectedRelicScRsp                  uint16 = 592
	LockRelicCsReq                             uint16 = 593
	RankUpEquipmentCsReq                       uint16 = 595
	GetRelicFilterPlanScRsp                    uint16 = 597
	ExchangeHcoinScRsp                         uint16 = 598
	AddRelicFilterPlanScRsp                    uint16 = 600
	PlayerSyncScNotify                         uint16 = 647
	GetStageLineupScRsp                        uint16 = 709
	SetLineupNameScRsp                         uint16 = 710
	JoinLineupScRsp                            uint16 = 711
	VirtualLineupTrialAvatarChangeScNotify     uint16 = 714
	GetCurLineupDataScRsp                      uint16 = 717
	QuitLineupScRsp                            uint16 = 720
	VirtualLineupDestroyNotify                 uint16 = 735
	SetLineupNameCsReq                         uint16 = 737
	GetCurLineupDataCsReq                      uint16 = 738
	GetLineupAvatarDataCsReq                   uint16 = 745
	ReplaceLineupScRsp                         uint16 = 746
	ChangeLineupLeaderScRsp                    uint16 = 750
	JoinLineupCsReq                            uint16 = 751
	ChangeLineupLeaderCsReq                    uint16 = 752
	SyncLineupNotify                           uint16 = 761
	GetAllLineupDataScRsp                      uint16 = 762
	GetAllLineupDataCsReq                      uint16 = 767
	GetLineupAvatarDataScRsp                   uint16 = 776
	SwitchLineupIndexCsReq                     uint16 = 781
	QuitLineupCsReq                            uint16 = 789
	SwapLineupScRsp                            uint16 = 791
	ExtraLineupDestroyNotify                   uint16 = 792
	SwitchLineupIndexScRsp                     uint16 = 793
	SwapLineupCsReq                            uint16 = 795
	ReplaceLineupCsReq                         uint16 = 798
	GetMailScRsp                               uint16 = 809
	DelMailScRsp                               uint16 = 811
	MarkReadMailScRsp                          uint16 = 817
	TakeMailAttachmentScRsp                    uint16 = 820
	MarkReadMailCsReq                          uint16 = 838
	GetMailCsReq                               uint16 = 847
	DelMailCsReq                               uint16 = 851
	TakeMailAttachmentCsReq                    uint16 = 889
	NewMailScNotify                            uint16 = 895
	GetQuestDataScRsp                          uint16 = 909
	TakeQuestRewardScRsp                       uint16 = 917
	GetQuestRecordScRsp                        uint16 = 920
	TakeQuestRewardCsReq                       uint16 = 938
	TakeQuestOptionalRewardCsReq               uint16 = 945
	GetQuestDataCsReq                          uint16 = 947
	FinishQuestScRsp                           uint16 = 961
	TakeQuestOptionalRewardScRsp               uint16 = 976
	BatchGetQuestDataScRsp                     uint16 = 981
	GetQuestRecordCsReq                        uint16 = 989
	FinishQuestCsReq                           uint16 = 991
	QuestRecordScNotify                        uint16 = 995
	PCLPIMJPADF                                uint16 = 1009
	ECHFCICIAGO                                uint16 = 1011
	NMFBEMKIBED                                uint16 = 1017
	KJEBHAEAPCG                                uint16 = 1020
	DJNJLMODHFB                                uint16 = 1038
	GCJNDOAGDIF                                uint16 = 1045
	OKBDGICAGAA                                uint16 = 1047
	MFCMNNJGHDJ                                uint16 = 1050
	NEDMNNMKCFD                                uint16 = 1051
	FOCOPOALNFP                                uint16 = 1052
	GPAFHGJOCGH                                uint16 = 1061
	BIAPIHNJHIN                                uint16 = 1076
	DNCIJBOACFO                                uint16 = 1089
	HHMDEJFGFKG                                uint16 = 1091
	AILPJKHCGGB                                uint16 = 1095
	PJMOBPLICEO                                uint16 = 1109
	HFICGFPKIJI                                uint16 = 1147
	GetMainMissionCustomValueCsReq             uint16 = 1201
	UpdateMainMissionCustomValueCsReq          uint16 = 1203
	UpdateTrackMainMissionIdScRsp              uint16 = 1206
	GetMissionDataScRsp                        uint16 = 1209
	SyncTaskCsReq                              uint16 = 1211
	AcceptMainMissionScRsp                     uint16 = 1213
	SubMissionRewardScNotify                   uint16 = 1214
	StartFinishSubMissionScNotify              uint16 = 1215
	FinishTalkMissionScRsp                     uint16 = 1217
	TeleportToMissionResetPointCsReq           uint16 = 1226
	DECPBAPAEGE                                uint16 = 1230
	AAEAIIFGMEE                                uint16 = 1233
	FinishTalkMissionCsReq                     uint16 = 1238
	MissionGroupWarnScNotify                   uint16 = 1245
	GetMissionDataCsReq                        uint16 = 1247
	GetMainMissionCustomValueScRsp             uint16 = 1249
	MissionRewardScNotify                      uint16 = 1251
	FinishCosumeItemMissionScRsp               uint16 = 1252
	UpdateMainMissionCustomValueScRsp          uint16 = 1255
	TeleportToMissionResetPointScRsp           uint16 = 1257
	GetMissionStatusScRsp                      uint16 = 1262
	GetMissionStatusCsReq                      uint16 = 1267
	UpdateTrackMainMissionIdCsReq              uint16 = 1269
	AcceptMainMissionCsReq                     uint16 = 1270
	StartFinishMainMissionScNotify             uint16 = 1273
	FinishCosumeItemMissionCsReq               uint16 = 1276
	MissionAcceptScNotify                      uint16 = 1277
	FinishedMissionScNotify                    uint16 = 1288
	SyncTaskScRsp                              uint16 = 1289
	EnterAdventureScRsp                        uint16 = 1309
	QuickStartCocoonStageScRsp                 uint16 = 1311
	GetFarmStageGachaInfoScRsp                 uint16 = 1317
	QuickStartFarmElementScRsp                 uint16 = 1320
	GetFarmStageGachaInfoCsReq                 uint16 = 1338
	FarmElementSweepScRsp                      uint16 = 1345
	QuickStartCocoonStageCsReq                 uint16 = 1351
	FarmElementSweepCsReq                      uint16 = 1361
	QuickStartFarmElementCsReq                 uint16 = 1389
	CocoonSweepScRsp                           uint16 = 1391
	CocoonSweepCsReq                           uint16 = 1395
	RecoverAllLineupCsReq                      uint16 = 1401
	ReEnterLastElementStageCsReq               uint16 = 1402
	SetClientPausedCsReq                       uint16 = 1403
	SceneGroupRefreshScNotify                  uint16 = 1404
	UnlockedAreaMapScNotify                    uint16 = 1405
	GetUnlockTeleportCsReq                     uint16 = 1406
	EnteredSceneChangeScNotify                 uint16 = 1408
	SceneEntityMoveScRsp                       uint16 = 1409
	SpringRefreshScRsp                         uint16 = 1410
	SceneCastSkillScRsp                        uint16 = 1411
	DeleteSummonUnitScRsp                      uint16 = 1412
	SetCurInteractEntityScRsp                  uint16 = 1413
	SceneReviveAfterRebattleScRsp              uint16 = 1414
	DeleteSummonUnitCsReq                      uint16 = 1416
	InteractPropScRsp                          uint16 = 1417
	ReEnterLastElementStageScRsp               uint16 = 1418
	StartCocoonStageCsReq                      uint16 = 1419
	GetCurSceneInfoScRsp                       uint16 = 1420
	GroupStateChangeCsReq                      uint16 = 1421
	OpenChestScNotify                          uint16 = 1422
	GameplayCounterCountDownCsReq              uint16 = 1423
	GameplayCounterRecoverCsReq                uint16 = 1424
	EnterSectionCsReq                          uint16 = 1426
	SetGroupCustomSaveDataScRsp                uint16 = 1427
	GetSceneMapInfoCsReq                       uint16 = 1428
	UpdateGroupPropertyScRsp                   uint16 = 1429
	DeactivateFarmElementCsReq                 uint16 = 1433
	ReturnLastTownScRsp                        uint16 = 1435
	ActivateFarmElementCsReq                   uint16 = 1436
	SpringRefreshCsReq                         uint16 = 1437
	InteractPropCsReq                          uint16 = 1438
	UpdateMechanismBarScNotify                 uint16 = 1439
	SyncServerSceneChangeNotify                uint16 = 1440
	UnlockTeleportNotify                       uint16 = 1442
	TrainWorldIdChangeScNotify                 uint16 = 1443
	GameplayCounterUpdateScNotify              uint16 = 1444
	SceneUpdatePositionVersionNotify           uint16 = 1445
	SceneEnterStageScRsp                       uint16 = 1446
	SceneEntityMoveCsReq                       uint16 = 1447
	EnterSceneByServerScNotify                 uint16 = 1448
	RecoverAllLineupScRsp                      uint16 = 1449
	SceneCastSkillCostMpScRsp                  uint16 = 1450
	SceneCastSkillCsReq                        uint16 = 1451
	SceneCastSkillCostMpCsReq                  uint16 = 1452
	EnterSceneCsReq                            uint16 = 1453
	GetSceneMapInfoScRsp                       uint16 = 1454
	SetClientPausedScRsp                       uint16 = 1455
	RefreshTriggerByClientScNotify             uint16 = 1456
	EnterSectionScRsp                          uint16 = 1457
	GroupStateChangeScRsp                      uint16 = 1458
	GetEnteredSceneScRsp                       uint16 = 1460
	SceneEntityMoveScNotify                    uint16 = 1461
	UpdateGroupPropertyCsReq                   uint16 = 1463
	RefreshTriggerByClientScRsp                uint16 = 1464
	DeactivateFarmElementScRsp                 uint16 = 1466
	LastSpringRefreshTimeNotify                uint16 = 1467
	StartCocoonStageScRsp                      uint16 = 1469
	GameplayCounterRecoverScRsp                uint16 = 1471
	RefreshTriggerByClientCsReq                uint16 = 1475
	SyncEntityBuffChangeListScNotify           uint16 = 1476
	SavePointsInfoNotify                       uint16 = 1477
	GetUnlockTeleportScRsp                     uint16 = 1478
	SceneEntityTeleportCsReq                   uint16 = 1480
	SceneCastSkillMpUpdateScNotify             uint16 = 1481
	SetTrainWorldIdScRsp                       uint16 = 1482
	SceneEntityTeleportScRsp                   uint16 = 1483
	SetTrainWorldIdCsReq                       uint16 = 1484
	ActivateFarmElementScRsp                   uint16 = 1485
	ScenePlaneEventScNotify                    uint16 = 1486
	GetEnteredSceneCsReq                       uint16 = 1487
	EntityBindPropScRsp                        uint16 = 1488
	GetCurSceneInfoCsReq                       uint16 = 1489
	GameplayCounterCountDownScRsp              uint16 = 1490
	ChangePropTimelineInfoScRsp                uint16 = 1491
	SceneReviveAfterRebattleCsReq              uint16 = 1492
	GroupStateChangeScNotify                   uint16 = 1494
	ChangePropTimelineInfoCsReq                uint16 = 1495
	EnterSceneScRsp                            uint16 = 1496
	SceneEnterStageCsReq                       uint16 = 1498
	UpdateFloorSavedValueNotify                uint16 = 1499
	GetShopListScRsp                           uint16 = 1509
	TakeCityShopRewardScRsp                    uint16 = 1511
	BuyGoodsScRsp                              uint16 = 1517
	BuyGoodsCsReq                              uint16 = 1538
	GetShopListCsReq                           uint16 = 1547
	TakeCityShopRewardCsReq                    uint16 = 1551
	CityShopInfoScNotify                       uint16 = 1589
	GetTutorialScRsp                           uint16 = 1609
	UnlockTutorialScRsp                        uint16 = 1611
	GetTutorialGuideScRsp                      uint16 = 1617
	UnlockTutorialGuideScRsp                   uint16 = 1620
	GetTutorialGuideCsReq                      uint16 = 1638
	FinishTutorialGuideScRsp                   uint16 = 1645
	GetTutorialCsReq                           uint16 = 1647
	UnlockTutorialCsReq                        uint16 = 1651
	FinishTutorialGuideCsReq                   uint16 = 1661
	UnlockTutorialGuideCsReq                   uint16 = 1689
	FinishTutorialScRsp                        uint16 = 1691
	FinishTutorialCsReq                        uint16 = 1695
	GetChallengeScRsp                          uint16 = 1709
	StartPartialChallengeCsReq                 uint16 = 1710
	LeaveChallengeScRsp                        uint16 = 1711
	ChallengeBossPhaseSettleNotify             uint16 = 1714
	StartChallengeScRsp                        uint16 = 1717
	RestartChallengePhaseCsReq                 uint16 = 1735
	GetChallengeGroupStatisticsScRsp           uint16 = 1737
	StartChallengeCsReq                        uint16 = 1738
	ChallengeLineupNotify                      uint16 = 1745
	EnterChallengeNextPhaseCsReq               uint16 = 1746
	GetChallengeCsReq                          uint16 = 1747
	TakeChallengeRewardCsReq                   uint16 = 1750
	LeaveChallengeCsReq                        uint16 = 1751
	GetCurChallengeScRsp                       uint16 = 1761
	StartPartialChallengeScRsp                 uint16 = 1767
	TakeChallengeRewardScRsp                   uint16 = 1781
	ChallengeSettleNotify                      uint16 = 1789
	GetCurChallengeCsReq                       uint16 = 1791
	EnterChallengeNextPhaseScRsp               uint16 = 1792
	GetChallengeGroupStatisticsCsReq           uint16 = 1793
	RestartChallengePhaseScRsp                 uint16 = 1798
	NOOHANIHDOB                                uint16 = 1801
	MILMMEFNLNJ                                uint16 = 1803
	BOBPIMEIKNI                                uint16 = 1808
	EJBIMMPCPIO                                uint16 = 1809
	NJKDJFMNKHG                                uint16 = 1810
	AAJKBEEBDAM                                uint16 = 1811
	HGILBKEAANN                                uint16 = 1813
	EBMIGHOFADN                                uint16 = 1814
	IKIOKKGFNEG                                uint16 = 1816
	LMBFKNBBIAN                                uint16 = 1817
	DNCMGMKKEHP                                uint16 = 1819
	EDEFMJMJHDG                                uint16 = 1820
	AHIFJEGEGBH                                uint16 = 1823
	JKGANDLBMDP                                uint16 = 1828
	IEJDCCKFGEC                                uint16 = 1830
	GHPMFEBLDAM                                uint16 = 1833
	AIMEOHGGGPH                                uint16 = 1835
	FLMKDFPNLNN                                uint16 = 1838
	DJHJNJIKKPI                                uint16 = 1839
	KAMDNPBKKEL                                uint16 = 1840
	IHCJAJGPJKK                                uint16 = 1842
	NGGMFNPNKEN                                uint16 = 1846
	HGJJMEKHJAB                                uint16 = 1847
	LOEBIPHOOHN                                uint16 = 1848
	NIJGLMIEFIG                                uint16 = 1850
	KGAECMDGBNH                                uint16 = 1851
	BKHJBHHFIBE                                uint16 = 1854
	DDEOCJJMGEA                                uint16 = 1855
	CELPBBFEKOP                                uint16 = 1856
	PFIFBAKBDJF                                uint16 = 1860
	LAFLMCHDLME                                uint16 = 1862
	LKHCFFGAEKF                                uint16 = 1863
	BFKHHONBNMB                                uint16 = 1864
	PIIIJLCPAML                                uint16 = 1865
	LOKBANIDPPG                                uint16 = 1867
	KDHBFOHEBPD                                uint16 = 1870
	IPFIDHAANHM                                uint16 = 1872
	EKCDCNHJIBC                                uint16 = 1874
	NNAFACGKDOB                                uint16 = 1875
	MBCOMIBJKID                                uint16 = 1877
	GMAHMMHHMAB                                uint16 = 1879
	IJFGBBJNDMP                                uint16 = 1881
	MJMMOIAINAK                                uint16 = 1882
	PMJFLCCJGNB                                uint16 = 1886
	JDIGBEDAFAK                                uint16 = 1887
	CFCMNGNHGFB                                uint16 = 1888
	FDDECGPJDKD                                uint16 = 1889
	JGMGPJBILHJ                                uint16 = 1890
	CBEDDKNMNDO                                uint16 = 1892
	DFEOPEHIBPM                                uint16 = 1893
	BODKPNEFBBP                                uint16 = 1894
	KBHGEGNHGPD                                uint16 = 1898
	DFNEHAIHGKK                                uint16 = 1899
	GetGachaInfoScRsp                          uint16 = 1909
	GetGachaCeilingScRsp                       uint16 = 1911
	DoGachaScRsp                               uint16 = 1917
	ExchangeGachaCeilingScRsp                  uint16 = 1920
	DoGachaCsReq                               uint16 = 1938
	GetGachaInfoCsReq                          uint16 = 1947
	GetGachaCeilingCsReq                       uint16 = 1951
	ExchangeGachaCeilingCsReq                  uint16 = 1989
	SetGachaDecideItemScRsp                    uint16 = 1991
	SetGachaDecideItemCsReq                    uint16 = 1995
	GetNpcTakenRewardScRsp                     uint16 = 2109
	GetFirstTalkNpcScRsp                       uint16 = 2111
	TakeTalkRewardScRsp                        uint16 = 2117
	FinishFirstTalkNpcScRsp                    uint16 = 2120
	TakeTalkRewardCsReq                        uint16 = 2138
	GetFirstTalkByPerformanceNpcScRsp          uint16 = 2145
	GetNpcTakenRewardCsReq                     uint16 = 2147
	GetFirstTalkNpcCsReq                       uint16 = 2151
	FinishFirstTalkByPerformanceNpcScRsp       uint16 = 2152
	GetFirstTalkByPerformanceNpcCsReq          uint16 = 2161
	FinishFirstTalkByPerformanceNpcCsReq       uint16 = 2176
	FinishFirstTalkNpcCsReq                    uint16 = 2189
	SelectInclinationTextScRsp                 uint16 = 2191
	SelectInclinationTextCsReq                 uint16 = 2195
	StartRaidScRsp                             uint16 = 2209
	DelSaveRaidScNotify                        uint16 = 2210
	LeaveRaidScRsp                             uint16 = 2217
	TakeChallengeRaidRewardCsReq               uint16 = 2220
	GetAllSaveRaidScRsp                        uint16 = 2237
	LeaveRaidCsReq                             uint16 = 2238
	GetRaidInfoScRsp                           uint16 = 2245
	StartRaidCsReq                             uint16 = 2247
	GetSaveRaidCsReq                           uint16 = 2250
	RaidInfoNotify                             uint16 = 2251
	SetClientRaidTargetCountScRsp              uint16 = 2252
	GetRaidInfoCsReq                           uint16 = 2261
	RaidKickByServerScNotify                   uint16 = 2267
	SetClientRaidTargetCountCsReq              uint16 = 2276
	GetSaveRaidScRsp                           uint16 = 2281
	GetChallengeRaidInfoScRsp                  uint16 = 2289
	ChallengeRaidNotify                        uint16 = 2291
	GetAllSaveRaidCsReq                        uint16 = 2293
	TakeChallengeRaidRewardScRsp               uint16 = 2295
	GetArchiveDataScRsp                        uint16 = 2309
	GetUpdatedArchiveDataScRsp                 uint16 = 2317
	GetUpdatedArchiveDataCsReq                 uint16 = 2338
	GetArchiveDataCsReq                        uint16 = 2347
	GetBigDataAllRecommendScRsp                uint16 = 2417
	GetBigDataRecommendScRsp                   uint16 = 2425
	GetBigDataRecommendCsReq                   uint16 = 2431
	GetBigDataAllRecommendCsReq                uint16 = 2446
	GetExpeditionDataScRsp                     uint16 = 2502
	AcceptExpeditionScRsp                      uint16 = 2505
	AcceptMultipleExpeditionScRsp              uint16 = 2507
	CancelActivityExpeditionCsReq              uint16 = 2509
	TakeMultipleExpeditionRewardScRsp          uint16 = 2510
	AcceptActivityExpeditionScRsp              uint16 = 2512
	CancelActivityExpeditionScRsp              uint16 = 2513
	AcceptActivityExpeditionCsReq              uint16 = 2517
	GetExpeditionDataCsReq                     uint16 = 2519
	TakeExpeditionRewardScRsp                  uint16 = 2525
	TakeExpeditionRewardCsReq                  uint16 = 2531
	CancelExpeditionScRsp                      uint16 = 2533
	TakeActivityExpeditionRewardScRsp          uint16 = 2534
	CancelExpeditionCsReq                      uint16 = 2535
	TakeMultipleActivityExpeditionRewardScRsp  uint16 = 2540
	TakeActivityExpeditionRewardCsReq          uint16 = 2543
	ExpeditionDataChangeScNotify               uint16 = 2546
	AcceptExpeditionCsReq                      uint16 = 2548
	TakeMultipleActivityExpeditionRewardCsReq  uint16 = 2550
	EGKKAPICIMP                                uint16 = 2551
	DAHLBDGMHCH                                uint16 = 2552
	DCAOKHOAHMD                                uint16 = 2555
	JMDPENNDDHA                                uint16 = 2557
	AJDKJFFJAME                                uint16 = 2559
	ADHBNIFDGLD                                uint16 = 2560
	ICIDMOINJOO                                uint16 = 2561
	LBLDGLJFLMC                                uint16 = 2564
	MDNGDCFHDGG                                uint16 = 2566
	FNLLDOPIPHD                                uint16 = 2567
	AIEAADPOFKA                                uint16 = 2569
	JGNEMNFKCGJ                                uint16 = 2574
	MAAANABLHCF                                uint16 = 2575
	LPHFJLKINIO                                uint16 = 2576
	FNIKJFLFMHP                                uint16 = 2577
	OMNDPBIJKEO                                uint16 = 2579
	EKAMFGAIBIB                                uint16 = 2581
	JKLCPKEJGNO                                uint16 = 2583
	LOFBEJOFPNB                                uint16 = 2586
	FNLGJFEOIMD                                uint16 = 2589
	CPOANNNHKGP                                uint16 = 2590
	KEBECGMKLNM                                uint16 = 2593
	HHMDIBHFIFG                                uint16 = 2594
	BKBJGLIOAJC                                uint16 = 2595
	JELFFFKNFCH                                uint16 = 2598
	GGLHJMHLPOI                                uint16 = 2600
	CurTrialActivityScNotify                   uint16 = 2602
	GetLoginActivityScRsp                      uint16 = 2609
	SubmitMaterialSubmitActivityMaterialCsReq  uint16 = 2610
	GetActivityScheduleConfigScRsp             uint16 = 2611
	TakeLoginActivityRewardScRsp               uint16 = 2617
	GetAvatarDeliverRewardActivityDataCsReq    uint16 = 2618
	StartTrialActivityScRsp                    uint16 = 2627
	LeaveTrialActivityCsReq                    uint16 = 2631
	GetTrialActivityDataCsReq                  uint16 = 2632
	TrialActivityDataChangeScNotify            uint16 = 2634
	TakeMaterialSubmitActivityRewardScRsp      uint16 = 2635
	GetMaterialSubmitActivityDataScRsp         uint16 = 2637
	TakeLoginActivityRewardCsReq               uint16 = 2638
	TakeTrialActivityRewardScRsp               uint16 = 2639
	LeaveTrialActivityScRsp                    uint16 = 2641
	GetLoginActivityCsReq                      uint16 = 2647
	AvatarDeliverRewardTakeRewardScRsp         uint16 = 2648
	GetActivityScheduleConfigCsReq             uint16 = 2651
	AvatarDeliverRewardChooseAvatarScRsp       uint16 = 2653
	TakeMaterialSubmitActivityRewardCsReq      uint16 = 2662
	TakeTrialActivityRewardCsReq               uint16 = 2665
	SubmitMaterialSubmitActivityMaterialScRsp  uint16 = 2667
	EnterTrialActivityStageScRsp               uint16 = 2672
	StartTrialActivityCsReq                    uint16 = 2679
	GetAvatarDeliverRewardActivityDataScRsp    uint16 = 2680
	AvatarDeliverRewardChooseAvatarCsReq       uint16 = 2683
	GetMaterialSubmitActivityDataCsReq         uint16 = 2693
	AvatarDeliverRewardTakeRewardCsReq         uint16 = 2696
	GetTrialActivityDataScRsp                  uint16 = 2700
	GetNpcMessageGroupScRsp                    uint16 = 2709
	FinishItemIdScRsp                          uint16 = 2711
	GetNpcStatusScRsp                          uint16 = 2717
	FinishSectionIdScRsp                       uint16 = 2720
	GetNpcStatusCsReq                          uint16 = 2738
	GetMissionMessageInfoScRsp                 uint16 = 2745
	GetNpcMessageGroupCsReq                    uint16 = 2747
	FinishItemIdCsReq                          uint16 = 2751
	GetMissionMessageInfoCsReq                 uint16 = 2761
	FinishSectionIdCsReq                       uint16 = 2789
	FinishPerformSectionIdScRsp                uint16 = 2791
	FinishPerformSectionIdCsReq                uint16 = 2795
	GetPlayerBoardDataScRsp                    uint16 = 2809
	SetDisplayAvatarScRsp                      uint16 = 2811
	SetHeadIconScRsp                           uint16 = 2817
	SetIsDisplayAvatarInfoScRsp                uint16 = 2820
	SetHeadIconCsReq                           uint16 = 2838
	SetAssistAvatarCsReq                       uint16 = 2845
	GetPlayerBoardDataCsReq                    uint16 = 2847
	SetPersonalCardScRsp                       uint16 = 2850
	SetDisplayAvatarCsReq                      uint16 = 2851
	SetPersonalCardCsReq                       uint16 = 2852
	SetSignatureScRsp                          uint16 = 2861
	SetAssistAvatarScRsp                       uint16 = 2876
	SetIsDisplayAvatarInfoCsReq                uint16 = 2889
	SetSignatureCsReq                          uint16 = 2891
	GetCurAssistCsReq                          uint16 = 2901
	GetFriendRecommendLineupDetailCsReq        uint16 = 2902
	CurAssistChangedNotify                     uint16 = 2903
	TakeAssistRewardCsReq                      uint16 = 2906
	SetFriendMarkScRsp                         uint16 = 2907
	GetFriendListInfoScRsp                     uint16 = 2909
	GetFriendRecommendListInfoCsReq            uint16 = 2910
	GetFriendApplyListInfoScRsp                uint16 = 2911
	SetAssistScRsp                             uint16 = 2913
	DeleteBlacklistScRsp                       uint16 = 2914
	GetAssistListCsReq                         uint16 = 2915
	GetPlayerDetailInfoScRsp                   uint16 = 2917
	GetFriendRecommendLineupDetailScRsp        uint16 = 2918
	GetAssistHistoryScRsp                      uint16 = 2919
	ApplyFriendScRsp                           uint16 = 2920
	SetFriendMarkCsReq                         uint16 = 2925
	SearchPlayerCsReq                          uint16 = 2926
	GetFriendDevelopmentInfoScRsp              uint16 = 2927
	GetPlatformPlayerInfoScRsp                 uint16 = 2930
	GetFriendRecommendLineupCsReq              uint16 = 2931
	GetFriendAssistListScRsp                   uint16 = 2932
	GetFriendLoginInfoCsReq                    uint16 = 2933
	SetFriendRemarkNameScRsp                   uint16 = 2935
	SetForbidOtherApplyFriendCsReq             uint16 = 2936
	SyncAddBlacklistScNotify                   uint16 = 2937
	GetPlayerDetailInfoCsReq                   uint16 = 2938
	GetFriendBattleRecordDetailScRsp           uint16 = 2939
	GetFriendRecommendLineupScRsp              uint16 = 2941
	SyncHandleFriendScNotify                   uint16 = 2945
	ReportPlayerScRsp                          uint16 = 2946
	GetFriendListInfoCsReq                     uint16 = 2947
	GetCurAssistScRsp                          uint16 = 2949
	SyncDeleteFriendScNotify                   uint16 = 2950
	GetFriendApplyListInfoCsReq                uint16 = 2951
	DeleteFriendScRsp                          uint16 = 2952
	GetPlatformPlayerInfoCsReq                 uint16 = 2955
	SearchPlayerScRsp                          uint16 = 2957
	HandleFriendScRsp                          uint16 = 2961
	SetFriendRemarkNameCsReq                   uint16 = 2962
	GetFriendBattleRecordDetailCsReq           uint16 = 2965
	GetFriendLoginInfoScRsp                    uint16 = 2966
	GetFriendRecommendListInfoScRsp            uint16 = 2967
	NewAssistHistoryNotify                     uint16 = 2969
	SetAssistCsReq                             uint16 = 2970
	GetAssistListScRsp                         uint16 = 2973
	DeleteFriendCsReq                          uint16 = 2976
	GetAssistHistoryCsReq                      uint16 = 2977
	GetFriendDevelopmentInfoCsReq              uint16 = 2979
	AddBlacklistCsReq                          uint16 = 2981
	SetForbidOtherApplyFriendScRsp             uint16 = 2985
	TakeAssistRewardScRsp                      uint16 = 2988
	ApplyFriendCsReq                           uint16 = 2989
	HandleFriendCsReq                          uint16 = 2991
	DeleteBlacklistCsReq                       uint16 = 2992
	AddBlacklistScRsp                          uint16 = 2993
	SyncApplyFriendScNotify                    uint16 = 2995
	GetFriendAssistListCsReq                   uint16 = 2997
	ReportPlayerCsReq                          uint16 = 2998
	BuyBpLevelCsReq                            uint16 = 3011
	TakeBpRewardCsReq                          uint16 = 3017
	BattlePassInfoNotify                       uint16 = 3047
	TakeBpRewardScRsp                          uint16 = 3051
	BuyBpLevelScRsp                            uint16 = 3089
	TakeAllRewardScRsp                         uint16 = 3095
	GetJukeboxDataScRsp                        uint16 = 3109
	UnlockBackGroundMusicScRsp                 uint16 = 3111
	PlayBackGroundMusicScRsp                   uint16 = 3117
	TrialBackGroundMusicScRsp                  uint16 = 3120
	PlayBackGroundMusicCsReq                   uint16 = 3138
	GetJukeboxDataCsReq                        uint16 = 3147
	TrialBackGroundMusicCsReq                  uint16 = 3189
	HHOEFAFIDKF                                uint16 = 3201
	GKEJEIJHOBM                                uint16 = 3209
	OFDJCKNCNJJ                                uint16 = 3210
	LKNDAIONPDK                                uint16 = 3211
	GOBECKLNBKI                                uint16 = 3213
	PBINMHDELNE                                uint16 = 3215
	DAKJPCHPDNE                                uint16 = 3217
	NFHONPNFGFC                                uint16 = 3220
	NGCKNKILNKH                                uint16 = 3238
	NFABOHBKKPM                                uint16 = 3246
	NKNILCNPOGJ                                uint16 = 3247
	HDFNBAECDLI                                uint16 = 3250
	FPLLNDGMCEL                                uint16 = 3251
	DMBFMIKOFJP                                uint16 = 3252
	DIJOAOKEOGN                                uint16 = 3257
	MHAHKDDGBMG                                uint16 = 3270
	BBALCNDAHAG                                uint16 = 3273
	DPLFMIDINFM                                uint16 = 3276
	DCPMEDCNFNI                                uint16 = 3281
	NMNBBAKGBCO                                uint16 = 3289
	MCALEHFBCNI                                uint16 = 3292
	NCJNACMOGDG                                uint16 = 3293
	GHBLGNMMCKJ                                uint16 = 3298
	TakeApRewardScRsp                          uint16 = 3309
	TakeAllApRewardCsReq                       uint16 = 3311
	GetDailyActiveInfoScRsp                    uint16 = 3317
	GetDailyActiveInfoCsReq                    uint16 = 3338
	TakeApRewardCsReq                          uint16 = 3347
	DailyActiveInfoNotify                      uint16 = 3351
	TakeAllApRewardScRsp                       uint16 = 3389
	GetRndOptionScRsp                          uint16 = 3409
	DailyFirstMeetPamScRsp                     uint16 = 3417
	DailyFirstMeetPamCsReq                     uint16 = 3438
	GetRndOptionCsReq                          uint16 = 3447
	CBNCNBMAHHG                                uint16 = 3509
	PICNBMNMICK                                uint16 = 3517
	IALBBBDEMEN                                uint16 = 3547
	GetFightActivityDataScRsp                  uint16 = 3609
	TakeFightActivityRewardCsReq               uint16 = 3611
	EnterFightActivityStageCsReq               uint16 = 3617
	FightActivityDataChangeScNotify            uint16 = 3638
	GetFightActivityDataCsReq                  uint16 = 3647
	EnterFightActivityStageScRsp               uint16 = 3651
	TakeFightActivityRewardScRsp               uint16 = 3689
	NMCKJKNBLKD                                uint16 = 3709
	CPKIJGOKEEI                                uint16 = 3711
	HODDGPADLPG                                uint16 = 3717
	CAFBJEBFAIC                                uint16 = 3720
	OOIGFJFMMHN                                uint16 = 3745
	JDLHPKNAEPH                                uint16 = 3747
	CJELMCNKCGN                                uint16 = 3751
	CMLJLFNKMBD                                uint16 = 3761
	ELMLHIMCCFL                                uint16 = 3789
	EPCJCLFDMPK                                uint16 = 3791
	AOPFAGDEFAN                                uint16 = 3795
	PALCBGGAIME                                uint16 = 3809
	DCHBPIPBGOG                                uint16 = 3811
	APANIBFAOED                                uint16 = 3817
	TextJoinQueryCsReq                         uint16 = 3838
	EDKHIIKKODJ                                uint16 = 3851
	SendMsgScRsp                               uint16 = 3909
	GetPrivateChatHistoryScRsp                 uint16 = 3911
	PrivateMsgOfflineUsersScNotify             uint16 = 3917
	GetChatFriendHistoryScRsp                  uint16 = 3920
	RevcMsgScNotify                            uint16 = 3938
	MarkChatEmojiScRsp                         uint16 = 3945
	SendMsgCsReq                               uint16 = 3947
	GetLoginChatInfoCsReq                      uint16 = 3950
	GetPrivateChatHistoryCsReq                 uint16 = 3951
	BatchMarkChatEmojiScRsp                    uint16 = 3952
	MarkChatEmojiCsReq                         uint16 = 3961
	GetLoginChatInfoScRsp                      uint16 = 3981
	GetChatFriendHistoryCsReq                  uint16 = 3989
	GetChatEmojiListScRsp                      uint16 = 3991
	GetChatEmojiListCsReq                      uint16 = 3995
	AcceptedPamMissionExpireScRsp              uint16 = 4009
	SyncAcceptedPamMissionNotify               uint16 = 4038
	AcceptedPamMissionExpireCsReq              uint16 = 4047
	MazeKillDirectCsReq                        uint16 = 4101
	SwitchMascotUpdateScNotify                 uint16 = 4103
	GetSwitchMascotDataCsReq                   uint16 = 4106
	ShareScRsp                                 uint16 = 4109
	TakePictureScRsp                           uint16 = 4111
	DifficultyAdjustmentUpdateDataScRsp        uint16 = 4113
	GetGunPlayDataScRsp                        uint16 = 4114
	DifficultyAdjustmentGetDataCsReq           uint16 = 4115
	GetShareDataScRsp                          uint16 = 4117
	CancelDirectDeliveryNoticeCsReq            uint16 = 4119
	UpdateGunPlayDataCsReq                     uint16 = 4126
	GetUnreleasedBlockInfoScRsp                uint16 = 4130
	GetMovieRacingDataScRsp                    uint16 = 4135
	GetOrigamiPropInfoScRsp                    uint16 = 4137
	GetShareDataCsReq                          uint16 = 4138
	SecurityReportScRsp                        uint16 = 4145
	UpdateMovieRacingDataScRsp                 uint16 = 4146
	ShareCsReq                                 uint16 = 4147
	MazeKillDirectScRsp                        uint16 = 4149
	SubmitOrigamiItemCsReq                     uint16 = 4150
	TakePictureCsReq                           uint16 = 4151
	TriggerVoiceScRsp                          uint16 = 4152
	GetUnreleasedBlockInfoCsReq                uint16 = 4155
	UpdateGunPlayDataScRsp                     uint16 = 4157
	GetMovieRacingDataCsReq                    uint16 = 4162
	CancelDirectDeliveryNoticeScRsp            uint16 = 4169
	DifficultyAdjustmentUpdateDataCsReq        uint16 = 4170
	DifficultyAdjustmentGetDataScRsp           uint16 = 4173
	TriggerVoiceCsReq                          uint16 = 4176
	DirectDeliveryScNotify                     uint16 = 4177
	SubmitOrigamiItemScRsp                     uint16 = 4181
	GetSwitchMascotDataScRsp                   uint16 = 4188
	CancelCacheNotifyScRsp                     uint16 = 4191
	GetGunPlayDataCsReq                        uint16 = 4192
	GetOrigamiPropInfoCsReq                    uint16 = 4193
	CancelCacheNotifyCsReq                     uint16 = 4195
	UpdateMovieRacingDataCsReq                 uint16 = 4198
	BPNMEBKDJEA                                uint16 = 4209
	KLNKDCOGKAB                                uint16 = 4211
	JCMIIGMBKED                                uint16 = 4217
	EMCHJNAOJML                                uint16 = 4220
	JFPLIKOOFEH                                uint16 = 4238
	OHFGLCDJBAJ                                uint16 = 4245
	DPCFNEKEAGF                                uint16 = 4247
	LJOKCKAGIPN                                uint16 = 4250
	GBOGGNAKMFG                                uint16 = 4251
	ILAFJJCACNA                                uint16 = 4252
	GLMFFPNGOFA                                uint16 = 4261
	IHCMMGAEBJL                                uint16 = 4276
	NDBJEMIIBLF                                uint16 = 4281
	MAJGIAAMOPL                                uint16 = 4289
	OPGOBPLGBOI                                uint16 = 4291
	BBPHBNHNNKJ                                uint16 = 4295
	LOELHDFPJKD                                uint16 = 4309
	DCDHMPBNLLA                                uint16 = 4310
	ABEEIIHOPKJ                                uint16 = 4311
	EOBDAMHNOCD                                uint16 = 4314
	JJBABPCHHFN                                uint16 = 4315
	ODIJHINADIN                                uint16 = 4317
	PBGPJKJIDJE                                uint16 = 4320
	LLGKDECGEHJ                                uint16 = 4326
	LJMDJDFODKM                                uint16 = 4335
	LKGGBAGPDNO                                uint16 = 4337
	OBABKHGHKFB                                uint16 = 4338
	NMEHFCJAACK                                uint16 = 4345
	EJPNNCNLONC                                uint16 = 4346
	LPMBCEGLLEN                                uint16 = 4347
	HDPCKNIPILG                                uint16 = 4350
	DELNPHOJHII                                uint16 = 4351
	FKIEEDNBPCO                                uint16 = 4352
	BCFIILJHALJ                                uint16 = 4357
	LCLEAOGMNML                                uint16 = 4361
	BCCMMOJGDPM                                uint16 = 4362
	AHJKNPINKIC                                uint16 = 4367
	NDGODKMMMFJ                                uint16 = 4376
	BMJHOMAFFKP                                uint16 = 4381
	AKMJJOPBNBL                                uint16 = 4389
	HOHHDIDEEMO                                uint16 = 4391
	EODCBBDGFDA                                uint16 = 4392
	IMLKDFCLBIL                                uint16 = 4393
	JNBBFOIODIA                                uint16 = 4395
	BFHHOIOBLMI                                uint16 = 4398
	TreasureDungeonFinishScNotify              uint16 = 4409
	InteractTreasureDungeonGridCsReq           uint16 = 4410
	UseTreasureDungeonItemScRsp                uint16 = 4435
	FightTreasureDungeonMonsterScRsp           uint16 = 4437
	GetTreasureDungeonActivityDataScRsp        uint16 = 4445
	QuitTreasureDungeonScRsp                   uint16 = 4446
	TreasureDungeonDataScNotify                uint16 = 4447
	OpenTreasureDungeonGridCsReq               uint16 = 4450
	EnterTreasureDungeonScRsp                  uint16 = 4452
	GetTreasureDungeonActivityDataCsReq        uint16 = 4461
	UseTreasureDungeonItemCsReq                uint16 = 4462
	InteractTreasureDungeonGridScRsp           uint16 = 4467
	EnterTreasureDungeonCsReq                  uint16 = 4476
	OpenTreasureDungeonGridScRsp               uint16 = 4481
	FightTreasureDungeonMonsterCsReq           uint16 = 4493
	QuitTreasureDungeonCsReq                   uint16 = 4498
	NGDGKNLHDBB                                uint16 = 4511
	EOENAEACHIG                                uint16 = 4517
	IPALIMDICEL                                uint16 = 4520
	BJJLCBBDHOF                                uint16 = 4538
	PBOMOMDNPNC                                uint16 = 4545
	DFHMNEKALKP                                uint16 = 4547
	IFNGOOHFDNF                                uint16 = 4551
	HMHOKLIPDEN                                uint16 = 4552
	ADLCNMFDMHH                                uint16 = 4561
	FNDIILLENPF                                uint16 = 4576
	BCIEDHOKLJN                                uint16 = 4589
	PEGNHMLBKJK                                uint16 = 4591
	JLLFEEOBFIB                                uint16 = 4595
	GAIEHAKNMMF                                uint16 = 4609
	CKGJHJGAEPP                                uint16 = 4611
	CNOOEFBFLCL                                uint16 = 4617
	MAMDABMCIMP                                uint16 = 4638
	PDPHCLPCNBP                                uint16 = 4647
	HPIDGCJDPEE                                uint16 = 4651
	GKFKFEJIPBL                                uint16 = 4703
	JGLPIMNCIOO                                uint16 = 4706
	GJDJCDBBHOE                                uint16 = 4709
	MCODBGAMGAE                                uint16 = 4710
	CEEICJCFGGC                                uint16 = 4713
	CJGHAPLBFAB                                uint16 = 4715
	NIBOCPDLGFG                                uint16 = 4717
	CJINJHFHOPB                                uint16 = 4719
	AEDKLLDFCPN                                uint16 = 4720
	DOFPCPJLEEC                                uint16 = 4726
	LCFIJCKNMDB                                uint16 = 4735
	HMLPKFEOHKH                                uint16 = 4745
	HEDNPLNCICC                                uint16 = 4747
	CHBOOADLNEA                                uint16 = 4749
	OKOJAAKFFDB                                uint16 = 4752
	EIKFIEJKMNC                                uint16 = 4755
	HOLMFMJBKBK                                uint16 = 4757
	HKNKCNMBEAO                                uint16 = 4761
	HJBDDADIPNL                                uint16 = 4762
	MGMDENGANGH                                uint16 = 4769
	OEODHMAFEBF                                uint16 = 4770
	FLEGGMNPHJD                                uint16 = 4773
	EIOPHCKBEEP                                uint16 = 4776
	NPEKKFMMKHL                                uint16 = 4777
	LIHOHDABGEC                                uint16 = 4788
	FBMPMKJCPKD                                uint16 = 4789
	DFJDECBJKMB                                uint16 = 4791
	KIGDOGNLBEJ                                uint16 = 4792
	PGHEIEEKELK                                uint16 = 4795
	HIFEIMICDML                                uint16 = 4798
	OICMOEJBNJD                                uint16 = 4801
	OJJADENFNDJ                                uint16 = 4802
	MJFPNIKJEFI                                uint16 = 4804
	ICEBBMAKHIE                                uint16 = 4805
	CHJJOFPENLA                                uint16 = 4807
	GetAetherDivideInfoScRsp                   uint16 = 4809
	HLBCKCDICAA                                uint16 = 4810
	FMHGDKEFBBJ                                uint16 = 4811
	GetAetherDivideInfoCsReq                   uint16 = 4812
	KEJCGNHOAML                                uint16 = 4813
	DBJKDLIOIEO                                uint16 = 4814
	MHJBACDKLFE                                uint16 = 4816
	LJNBHDOAOOC                                uint16 = 4820
	MHCGODLNNPJ                                uint16 = 4823
	LIHKFBKOGNB                                uint16 = 4824
	CJCJEFBPEBC                                uint16 = 4825
	ODIGIPEHBGG                                uint16 = 4826
	CMAFKHKKDNM                                uint16 = 4827
	ANDACJGGBEB                                uint16 = 4829
	NOFNMHPMCJN                                uint16 = 4830
	KNCCOBPCFEL                                uint16 = 4831
	AEINLINFCDL                                uint16 = 4833
	OPEKCNGFDJO                                uint16 = 4834
	MPPINLPANKB                                uint16 = 4835
	GEJIONMNIKJ                                uint16 = 4836
	EDAGAAECHLK                                uint16 = 4839
	FOJNPAEENOP                                uint16 = 4840
	LAFNPEBLNDE                                uint16 = 4843
	DMLLGODHJPA                                uint16 = 4844
	HEFCPAFNHHK                                uint16 = 4845
	DNHEIONHEIE                                uint16 = 4847
	DIOHKAKCKLC                                uint16 = 4848
	KGJPEBCEPDL                                uint16 = 4850
	GetBenefitActivityInfoScRsp                uint16 = 4852
	TakeBenefitActivityRewardScRsp             uint16 = 4855
	GetBenefitActivityInfoCsReq                uint16 = 4869
	JoinBenefitActivityScRsp                   uint16 = 4883
	JoinBenefitActivityCsReq                   uint16 = 4885
	TakeBenefitActivityRewardCsReq             uint16 = 4898
	AHGMKOGAHGD                                uint16 = 4909
	HEHLNIMBFPM                                uint16 = 4911
	AFGANAFOOIJ                                uint16 = 4917
	NKGDPMONOPD                                uint16 = 4938
	CEDICAHDICO                                uint16 = 4947
	PKIEOOOHPKG                                uint16 = 4951
	GetPhoneDataScRsp                          uint16 = 5109
	SelectPhoneThemeCsReq                      uint16 = 5111
	SelectChatBubbleScRsp                      uint16 = 5117
	UnlockPhoneThemeScNotify                   uint16 = 5120
	SelectChatBubbleCsReq                      uint16 = 5138
	GetPhoneDataCsReq                          uint16 = 5147
	UnlockChatBubbleScNotify                   uint16 = 5151
	UnlockPhoneCaseScNotify                    uint16 = 5161
	SelectPhoneThemeScRsp                      uint16 = 5189
	SelectPhoneCaseScRsp                       uint16 = 5191
	SelectPhoneCaseCsReq                       uint16 = 5195
	CCKKHNGHAJM                                uint16 = 5317
	KNCFBDFLDFE                                uint16 = 5320
	FFOOFOHOOBD                                uint16 = 5338
	JNAKJHBCEGK                                uint16 = 5351
	PFLKCKJCDNP                                uint16 = 5391
	OLJNKMLOAAI                                uint16 = 5395
	KIDIMAGODDO                                uint16 = 5402
	LILAOOMJGBO                                uint16 = 5403
	HLMLPPGFCOK                                uint16 = 5404
	OOAKNKPPFJN                                uint16 = 5406
	OKIECEJGBNH                                uint16 = 5425
	CCAMIJODAEO                                uint16 = 5427
	MJKHAJMOFHD                                uint16 = 5429
	HGPGNDFJFNC                                uint16 = 5431
	FNKBMEGFCAO                                uint16 = 5433
	NGOPJFJGHPF                                uint16 = 5434
	DIPABHGIOHC                                uint16 = 5435
	JOGMAKFDDDI                                uint16 = 5440
	KJLLINOKEEG                                uint16 = 5451
	ICHJABJDILN                                uint16 = 5452
	IKLEMEDKBFA                                uint16 = 5453
	GAAKNMCLFDB                                uint16 = 5454
	CEEEEKKFEEN                                uint16 = 5455
	HCIMDMNDAJH                                uint16 = 5459
	PDLHODMBDGG                                uint16 = 5460
	FDPNNADDHFB                                uint16 = 5465
	BJCGPIDDKCO                                uint16 = 5466
	HJPPOLJOBDC                                uint16 = 5468
	DCEINFONEKB                                uint16 = 5472
	AJIDLGFBLBC                                uint16 = 5477
	LDJCMNILFIG                                uint16 = 5478
	HJMLHAEOOGL                                uint16 = 5479
	HDELCMIKKHE                                uint16 = 5482
	HFGIKBFCNHJ                                uint16 = 5490
	OHKOBJKMJIF                                uint16 = 5492
	AKEADAIMKKB                                uint16 = 5497
	KODBMIFICKL                                uint16 = 5499
	BGEPEAIEKBL                                uint16 = 5502
	MDFHEJDFPPF                                uint16 = 5504
	EMGANNEHNLJ                                uint16 = 5506
	PJBEEHEFGII                                uint16 = 5508
	HJFPNHKKGEM                                uint16 = 5514
	BIGOGDODAFH                                uint16 = 5515
	MIGIKKGEHAB                                uint16 = 5525
	AEJFOFIFGJJ                                uint16 = 5527
	EKKDJDGIOCN                                uint16 = 5529
	FGBHBCKPEEC                                uint16 = 5530
	FMDBHHLPJLF                                uint16 = 5532
	DKMLFPLIOLM                                uint16 = 5535
	NCHANEEDHEB                                uint16 = 5536
	OCAGICPAPNN                                uint16 = 5538
	LEGNFIJHFJA                                uint16 = 5540
	IJIFCKAIFAM                                uint16 = 5541
	EGOEBBEKOAK                                uint16 = 5543
	AFBHPGLBLCH                                uint16 = 5545
	CEJFNFDKADC                                uint16 = 5546
	PBJPGCCKDAA                                uint16 = 5549
	IFNPPMEGAGK                                uint16 = 5552
	LOKEPEOLDLK                                uint16 = 5555
	GFHPIELAEMD                                uint16 = 5559
	BOPIBEAABHM                                uint16 = 5562
	NIPHNIOEOOH                                uint16 = 5565
	MJGELCNMLIM                                uint16 = 5566
	KKEJMJMKLBH                                uint16 = 5569
	LIHDJLDPBGO                                uint16 = 5571
	PHPGHELOJON                                uint16 = 5572
	MGCOCDHLOBL                                uint16 = 5573
	HJBINCHFMIJ                                uint16 = 5575
	ODOGHFDOAME                                uint16 = 5576
	EBCGENEHPLO                                uint16 = 5577
	NEEPIDOKIOE                                uint16 = 5579
	FOLENGGPDHF                                uint16 = 5580
	BMEMEAIPICO                                uint16 = 5584
	JOLMFHDNMJC                                uint16 = 5585
	AGMNCHOFGEP                                uint16 = 5588
	DHKBDLAKCJC                                uint16 = 5591
	OHPFCGMAOEK                                uint16 = 5592
	NJNDMAKLHGH                                uint16 = 5594
	LAMOKNACCFE                                uint16 = 5595
	AIBMGOGLFLJ                                uint16 = 5596
	CMJMFIKAJBG                                uint16 = 5597
	NHMDADNCDBN                                uint16 = 5598
	NDMBKKGADNO                                uint16 = 5602
	JMMJGOOJBEH                                uint16 = 5603
	CFIEHABLDOC                                uint16 = 5605
	AEHCIIHAMBJ                                uint16 = 5606
	BADNPCNJGOM                                uint16 = 5607
	ODFBPOALLPD                                uint16 = 5609
	LNDHKMMPBFJ                                uint16 = 5611
	HCEEKIOAEKN                                uint16 = 5612
	JKOFLFFJDBM                                uint16 = 5614
	NELFLKMAAAG                                uint16 = 5616
	FJGGNHODADK                                uint16 = 5617
	FFEBCHCIACF                                uint16 = 5618
	MPJCFEFLGJK                                uint16 = 5620
	HPEJDAODMLN                                uint16 = 5621
	JBENCNEAKPE                                uint16 = 5625
	EMPIBLJGIOC                                uint16 = 5627
	LILKJLIGCEL                                uint16 = 5628
	MKKJIHPPCNF                                uint16 = 5630
	CCBKNLDOLJH                                uint16 = 5631
	KCOMNODHCOD                                uint16 = 5633
	CADKIDNGMCB                                uint16 = 5634
	HNEFKBODKKN                                uint16 = 5636
	IJNIPLKKIBO                                uint16 = 5638
	LBOFIPJNCAP                                uint16 = 5639
	DELPONBAIFF                                uint16 = 5640
	EAMNGABAKGB                                uint16 = 5641
	BCHGLIJBDHA                                uint16 = 5643
	MOCFDOHFPCF                                uint16 = 5645
	EMDBEGPFMAO                                uint16 = 5647
	PMONFIOMJOC                                uint16 = 5648
	MLPGILOAGLH                                uint16 = 5650
	NJDDMAPMNKL                                uint16 = 5651
	NGPONCGIHAG                                uint16 = 5652
	OEHGLLPFDFA                                uint16 = 5654
	JCIGNGGPCBN                                uint16 = 5655
	HIBPKPMLBBP                                uint16 = 5656
	OJBDHNMJPLH                                uint16 = 5658
	KENLCMJBHMC                                uint16 = 5659
	MPLLHJCDKFN                                uint16 = 5661
	CBAFJJELCOI                                uint16 = 5665
	IBDFBFGGGHM                                uint16 = 5666
	GLFANALPNBG                                uint16 = 5668
	FHKEEDOLOBO                                uint16 = 5669
	BJDHAPHDEHH                                uint16 = 5672
	DAMECIJDLOE                                uint16 = 5676
	CFNGEEOFONA                                uint16 = 5678
	ICPJOFDGLPC                                uint16 = 5679
	AFKBBNDLPBE                                uint16 = 5680
	AEBNKGPCHPC                                uint16 = 5683
	MBNABDOJJJN                                uint16 = 5684
	EMILPJEHKKD                                uint16 = 5685
	NAHGHNMDPKB                                uint16 = 5686
	AOMJBJFEEAA                                uint16 = 5687
	JIOFFGNHLLM                                uint16 = 5688
	AGLBFJIIDKJ                                uint16 = 5689
	DIMMGDFHAED                                uint16 = 5691
	EKMBHACDNCN                                uint16 = 5692
	ANLNAJLMEOE                                uint16 = 5694
	BJAJDPCGDAE                                uint16 = 5695
	GGPBCFDAHGK                                uint16 = 5696
	GetBattleCollegeDataScRsp                  uint16 = 5709
	StartBattleCollegeCsReq                    uint16 = 5717
	BattleCollegeDataChangeScNotify            uint16 = 5738
	GetBattleCollegeDataCsReq                  uint16 = 5747
	StartBattleCollegeScRsp                    uint16 = 5751
	HeliobusActivityDataScRsp                  uint16 = 5809
	NGLFFFNCJLN                                uint16 = 5811
	LLMJHIHLBAD                                uint16 = 5817
	LFDIHDCPJCP                                uint16 = 5820
	ODELNDKFBKD                                uint16 = 5835
	FHDCAFMEFLJ                                uint16 = 5838
	MNILINDCMGJ                                uint16 = 5845
	IJCKOIFEHGO                                uint16 = 5846
	HeliobusActivityDataCsReq                  uint16 = 5847
	LGBFJCIBOIK                                uint16 = 5850
	KMHCOMICEIM                                uint16 = 5851
	LCFNJLBHFBA                                uint16 = 5852
	CNFNJCECHGO                                uint16 = 5861
	LCMIMFDLDIH                                uint16 = 5862
	HEONENLGLFB                                uint16 = 5867
	IDAEPKPFBEA                                uint16 = 5876
	DGMALFEMGCJ                                uint16 = 5889
	EIFABALINBA                                uint16 = 5891
	OMMHMMGBOME                                uint16 = 5892
	LHENFPKMHMK                                uint16 = 5893
	LJFMPHJCKBG                                uint16 = 5895
	CDLFDJHKNLA                                uint16 = 5898
	DEMGAOOOFGJ                                uint16 = 5902
	IDOGDCHCOMJ                                uint16 = 5905
	JDFGHDDLNJE                                uint16 = 5919
	NPGOMPIDMFJ                                uint16 = 5933
	DMFFKCBGKKJ                                uint16 = 5952
	CNACENHJLOL                                uint16 = 5955
	NANEFGENLAO                                uint16 = 5957
	HPAFBMDJODD                                uint16 = 5959
	MICGLCNILOL                                uint16 = 5960
	EGMPIICELAH                                uint16 = 5962
	MDLDAPDJNMK                                uint16 = 5963
	DPOLNPLHIAA                                uint16 = 5964
	JLKBPPLJKIJ                                uint16 = 5966
	MMPJPKEODEL                                uint16 = 5967
	EBCACFHPPCF                                uint16 = 5969
	EMOPJEGPHGK                                uint16 = 5970
	GHCCOIMOGJB                                uint16 = 5973
	NOEBCPDGJAC                                uint16 = 5974
	PANODPAJMOJ                                uint16 = 5975
	PJCCJAFEPMP                                uint16 = 5976
	OIFFDIOMIIF                                uint16 = 5977
	JDBGCNBDLLO                                uint16 = 5981
	PAMKNHOKHDL                                uint16 = 5983
	ADMNIBIJHFL                                uint16 = 5984
	HJFJAAFODAC                                uint16 = 5985
	FDCNFHLJEJH                                uint16 = 5986
	FJFKPKDAKGF                                uint16 = 5990
	FMLLCFAEGKF                                uint16 = 5993
	DHKKJDJOKMH                                uint16 = 5996
	PKKGGNJICEO                                uint16 = 5998
	CJMDEHEGBDB                                uint16 = 6000
	NDBCPOEKMAP                                uint16 = 6001
	CCFKEPHJECI                                uint16 = 6002
	AACAGDMNDIA                                uint16 = 6003
	CKNFINDGJON                                uint16 = 6005
	PGPKFJMBNBG                                uint16 = 6006
	JNAFHOCCEOA                                uint16 = 6007
	MKIBJDIGCOJ                                uint16 = 6008
	DDJJOEOEAFE                                uint16 = 6009
	JAACIPEBEFB                                uint16 = 6010
	GGBHOHNPHBO                                uint16 = 6011
	PCFFAIPODFI                                uint16 = 6012
	JBAGBPPIHFC                                uint16 = 6013
	IHAKOOAGOHF                                uint16 = 6014
	JBJBLDDKCFF                                uint16 = 6016
	HLBDHKLBFKL                                uint16 = 6017
	IKAIOFPBHJM                                uint16 = 6019
	LONGMLLNJHJ                                uint16 = 6020
	AOOCNHIDGNN                                uint16 = 6021
	CEMEPHAFBEN                                uint16 = 6023
	PCHHLGNHFJL                                uint16 = 6024
	OGBGKDNMJJN                                uint16 = 6025
	GOPCELBLMEK                                uint16 = 6027
	OMMPOGNONLL                                uint16 = 6028
	NLEOJELEADF                                uint16 = 6029
	OHCIMEBNOOA                                uint16 = 6030
	BNLFGHMPGLL                                uint16 = 6032
	BEBCPIBJNOI                                uint16 = 6033
	HPBOKBBJOHD                                uint16 = 6034
	KFAPGHFPOAM                                uint16 = 6035
	KPFIALIIHLK                                uint16 = 6036
	GIEMCFEHCII                                uint16 = 6037
	KCEJGCBBPPE                                uint16 = 6038
	OMLIBKJMKKC                                uint16 = 6040
	FPBFPJJHEAA                                uint16 = 6041
	OAKPNPEKJFF                                uint16 = 6042
	HIECGAOJIHM                                uint16 = 6043
	DNJFHOHBEIH                                uint16 = 6044
	CMMAFJGOEBO                                uint16 = 6045
	DGJOJAEFENJ                                uint16 = 6046
	LLOHGKBAPOE                                uint16 = 6047
	KBLGBOKEIJE                                uint16 = 6048
	RogueTournGetCurRogueCocoonInfoCsReq       uint16 = 6049
	FCHBLIEKDIO                                uint16 = 6050
	MMBHMGMADNP                                uint16 = 6051
	AKBJCJBIHLC                                uint16 = 6054
	PIBAKIJCEOL                                uint16 = 6055
	CFHDDAAJFIO                                uint16 = 6056
	OJLEHAPCALO                                uint16 = 6057
	PAOBPIIDBDB                                uint16 = 6058
	KLPKPGEPOPL                                uint16 = 6059
	EFKCPPCLINI                                uint16 = 6060
	OOACPALGDLJ                                uint16 = 6061
	DLIFNCLEKHL                                uint16 = 6062
	MDNBEPPCHCJ                                uint16 = 6063
	GBKEEFDDELM                                uint16 = 6064
	OOILNIEOEHE                                uint16 = 6065
	BNGDJIDDHNF                                uint16 = 6067
	LNLDMFOFFNM                                uint16 = 6069
	LKCMNAAIMMM                                uint16 = 6071
	BOKJMAGOFGA                                uint16 = 6072
	JIFGEAGMFGG                                uint16 = 6075
	MLMDECIBOEB                                uint16 = 6076
	HJPIIBCMGGC                                uint16 = 6077
	KPBBMANACIC                                uint16 = 6078
	RogueTournGetCurRogueCocoonInfoScRsp       uint16 = 6079
	MBAJJFBEGMI                                uint16 = 6080
	KAADNHDENPI                                uint16 = 6082
	KAPJNAKMKOB                                uint16 = 6083
	BFDADNAEHBG                                uint16 = 6086
	EABGBGMMPKN                                uint16 = 6088
	PDBEEBMJJDB                                uint16 = 6089
	DPFFEELFABM                                uint16 = 6090
	CCCJDJCBJLO                                uint16 = 6091
	HPLHFAIIDEO                                uint16 = 6093
	IEGIMEHHLKH                                uint16 = 6096
	DDPEFJCMLJL                                uint16 = 6097
	IDJKKNKJGFP                                uint16 = 6098
	BOAOLFBJAGJ                                uint16 = 6099
	OFHDPLMJPCA                                uint16 = 6100
	GetAllServerPrefsDataScRsp                 uint16 = 6109
	UpdateServerPrefsDataScRsp                 uint16 = 6111
	GetServerPrefsDataScRsp                    uint16 = 6117
	GetAllServerPrefsDataCsReq                 uint16 = 6147
	UpdateServerPrefsDataCsReq                 uint16 = 6151
	GetStoryLineInfoScRsp                      uint16 = 6209
	ChangeStoryLineFinishScNotify              uint16 = 6211
	StoryLineInfoScNotify                      uint16 = 6238
	GetStoryLineInfoCsReq                      uint16 = 6247
	StoryLineTrialAvatarChangeScNotify         uint16 = 6289
	GetHeartDialInfoScRsp                      uint16 = 6309
	SubmitEmotionItemScRsp                     uint16 = 6311
	ChangeScriptEmotionScRsp                   uint16 = 6317
	FinishEmotionDialoguePerformanceScRsp      uint16 = 6320
	ChangeScriptEmotionCsReq                   uint16 = 6338
	GetHeartDialInfoCsReq                      uint16 = 6347
	HeartDialTraceScriptScRsp                  uint16 = 6361
	FinishEmotionDialoguePerformanceCsReq      uint16 = 6389
	HeartDialTraceScriptCsReq                  uint16 = 6391
	HeartDialScriptChangeScNotify              uint16 = 6395
	ODNHKDDNLPB                                uint16 = 6409
	KALFPKOLCMC                                uint16 = 6410
	HELMGHHLLMB                                uint16 = 6411
	MBCONMAHPGC                                uint16 = 6420
	PDNGJBANLEH                                uint16 = 6435
	IEOHIBMALDB                                uint16 = 6437
	AKOOJPGGIID                                uint16 = 6438
	PIIHEFOPBGL                                uint16 = 6445
	LOEJLDBMELC                                uint16 = 6447
	GMOJKLOLCMI                                uint16 = 6450
	KNLCKNHNEIO                                uint16 = 6451
	LJAEKJBGAJH                                uint16 = 6461
	LJGHHMAFMJC                                uint16 = 6462
	PDFNJCMJFFE                                uint16 = 6467
	AADDJICOPGG                                uint16 = 6476
	HPNFJEGCFDM                                uint16 = 6481
	CPBMOIIPKMP                                uint16 = 6491
	GNGFHKPCGJO                                uint16 = 6493
	BFPILCMLFDN                                uint16 = 6495
	KDAAEBDKEJG                                uint16 = 6502
	JHCOIGILILA                                uint16 = 6505
	GCAOGCEIAFP                                uint16 = 6507
	JJIBHDBEHJE                                uint16 = 6510
	ABOADJAAGBL                                uint16 = 6512
	DPOGEMAPOKL                                uint16 = 6513
	BHHHPJOIJPH                                uint16 = 6517
	DHFHJJKPPPD                                uint16 = 6519
	ICIEJEBBKJE                                uint16 = 6520
	MCHKEHCFKKO                                uint16 = 6525
	CBIKDGJILJF                                uint16 = 6531
	AOCPKLPBLEF                                uint16 = 6533
	CKGENIOHBPL                                uint16 = 6534
	KDOMNBDDOBE                                uint16 = 6535
	HBDNGKNAKBL                                uint16 = 6543
	LBOLCAABCCG                                uint16 = 6546
	JOABMHCDHDL                                uint16 = 6548
	EnterEraFlipperRegionCsReq                 uint16 = 6555
	OJOBJIHLCAK                                uint16 = 6557
	BDAGNLNGHIC                                uint16 = 6559
	ChangeEraFlipperDataScRsp                  uint16 = 6562
	PKBPGMGMDAC                                uint16 = 6567
	GetEraFlipperDataScRsp                     uint16 = 6568
	AKHJLLNNAJA                                uint16 = 6570
	IHEOCPCAMBF                                uint16 = 6572
	BGDHCMBNNFC                                uint16 = 6574
	KGIPBNHOJLG                                uint16 = 6582
	CCBCMAOCEDG                                uint16 = 6584
	OCPJHDOFIMC                                uint16 = 6592
	PFFDLKILHFM                                uint16 = 6595
	CHJANNPMLEO                                uint16 = 6599
	LHBCEGEFIAC                                uint16 = 6609
	LFDJEIEKICL                                uint16 = 6617
	MAFFBCDOIEM                                uint16 = 6638
	JEEHLBCGFJK                                uint16 = 6647
	CDIEDODAHDO                                uint16 = 6651
	MFPFDOKCIPH                                uint16 = 6709
	KFNDBFAEMJN                                uint16 = 6711
	OHHPBOEJLLC                                uint16 = 6717
	DAGIOJOBGEE                                uint16 = 6720
	NFELMOCJCMM                                uint16 = 6738
	MIBIMKEPCME                                uint16 = 6745
	KJIEOKEFPAF                                uint16 = 6747
	LOPHALHNDHF                                uint16 = 6750
	PIAMJEHGGHH                                uint16 = 6751
	EGOCFONOBKF                                uint16 = 6752
	LJAEKBDJJFM                                uint16 = 6761
	DNBHCGKJKAH                                uint16 = 6776
	OGHNACAMIEC                                uint16 = 6789
	GBBMKCNJBOM                                uint16 = 6791
	PFLIEDNIGPA                                uint16 = 6795
	EnterMapRotationRegionScRsp                uint16 = 6809
	RemoveRotaterScRsp                         uint16 = 6810
	DeployRotaterScRsp                         uint16 = 6811
	InteractChargerScRsp                       uint16 = 6817
	RotateMapScRsp                             uint16 = 6820
	RemoveRotaterCsReq                         uint16 = 6837
	InteractChargerCsReq                       uint16 = 6838
	GetMapRotationDataScRsp                    uint16 = 6845
	EnterMapRotationRegionCsReq                uint16 = 6847
	LeaveMapRotationRegionScNotify             uint16 = 6850
	DeployRotaterCsReq                         uint16 = 6851
	ResetMapRotationRegionScRsp                uint16 = 6852
	GetMapRotationDataCsReq                    uint16 = 6861
	UpdateRotaterScNotify                      uint16 = 6867
	ResetMapRotationRegionCsReq                uint16 = 6876
	UpdateEnergyScNotify                       uint16 = 6881
	RotateMapCsReq                             uint16 = 6889
	LeaveMapRotationRegionScRsp                uint16 = 6891
	UpdateMapRotationDataScNotify              uint16 = 6893
	LeaveMapRotationRegionCsReq                uint16 = 6895
	TakeRollShopRewardScRsp                    uint16 = 6904
	DoGachaInRollShopScRsp                     uint16 = 6905
	GetRollShopInfoScRsp                       uint16 = 6906
	DoGachaInRollShopCsReq                     uint16 = 6909
	GetRollShopInfoCsReq                       uint16 = 6913
	TakeRollShopRewardCsReq                    uint16 = 6914
	OfferingInfoScNotify                       uint16 = 6923
	TakeOfferingRewardScRsp                    uint16 = 6924
	SubmitOfferingItemScRsp                    uint16 = 6925
	GetOfferingInfoScRsp                       uint16 = 6926
	SubmitOfferingItemCsReq                    uint16 = 6929
	GetOfferingInfoCsReq                       uint16 = 6933
	TakeOfferingRewardCsReq                    uint16 = 6934
	FBNKKFNMLOM                                uint16 = 6945
	RaidCollectionDataScRsp                    uint16 = 6946
	OOIBELOGDIA                                uint16 = 6949
	RaidCollectionDataCsReq                    uint16 = 6953
	IIKBHNKPJFO                                uint16 = 6954
	BLHIJCEDNNE                                uint16 = 6964
	JEPAGHNAPCC                                uint16 = 6965
	GALFJGOGDHL                                uint16 = 6966
	IEEJNPGCEGA                                uint16 = 6969
	LJDGJLPLHIH                                uint16 = 6973
	JPIKKKDADND                                uint16 = 6974
	JADFJJNOAAB                                uint16 = 6981
	KADIILDGAPD                                uint16 = 6982
	NPOPGIHLDNO                                uint16 = 6983
	AKLPOCNEDDJ                                uint16 = 6984
	MDDJJLJFACE                                uint16 = 6985
	KMCHEILEKKI                                uint16 = 6986
	OPHHNPGIKNK                                uint16 = 6987
	OGDMBJBOBDN                                uint16 = 6989
	LFHDJPDIHDO                                uint16 = 6990
	MFIEPPHJHBC                                uint16 = 6991
	FEPPFECHKPB                                uint16 = 6992
	IMMCALOKIJN                                uint16 = 6993
	NLBILPPLBFL                                uint16 = 6994
	HAIBDDPANJJ                                uint16 = 6995
	GBNCOMHBFJA                                uint16 = 6997
	IAACILJIODI                                uint16 = 6998
	FPKLBDNAPEG                                uint16 = 6999
	MGMACJHPELD                                uint16 = 7000
	JFGJMCCOPJF                                uint16 = 7001
	KGHHKLNJANA                                uint16 = 7002
	CKCIPGPKAHB                                uint16 = 7003
	NJPCCBNPMEP                                uint16 = 7004
	DLEOPIFLMAO                                uint16 = 7006
	MEKBAOOIFOK                                uint16 = 7008
	GMICFOLDOPA                                uint16 = 7009
	HJDFOHNJPPH                                uint16 = 7010
	FHLDKLFKEPG                                uint16 = 7011
	AGJIJJBBDKG                                uint16 = 7013
	JFONBMAAFEK                                uint16 = 7014
	CILGFNEGEHD                                uint16 = 7015
	LJIDIFFPAKL                                uint16 = 7017
	JFEEFPCMBPK                                uint16 = 7018
	JMFGGMFHEHJ                                uint16 = 7019
	HILFAIHNGKD                                uint16 = 7020
	DGEDOPCNLME                                uint16 = 7021
	HNPCLJOKJII                                uint16 = 7022
	ENEAMLFNHPE                                uint16 = 7023
	CCANODEDLHB                                uint16 = 7024
	OKGDEJIGADE                                uint16 = 7026
	EJPBCEIBNJI                                uint16 = 7027
	JFJGHAFDFJE                                uint16 = 7028
	GCCMLJNDNFH                                uint16 = 7030
	FJOBKKKDMNK                                uint16 = 7031
	GDKGEAJPECK                                uint16 = 7032
	LJPJGDEGOMH                                uint16 = 7033
	CMDMGBPBAID                                uint16 = 7034
	EKIOMGCIINK                                uint16 = 7035
	GDBCMFABPBJ                                uint16 = 7036
	NOHGHMMFBLC                                uint16 = 7037
	MKLFMDBBAON                                uint16 = 7038
	MMJEHAILNGE                                uint16 = 7039
	MAOFJEEAPJO                                uint16 = 7040
	IFEBFONCHCO                                uint16 = 7041
	FGOIFCABMMG                                uint16 = 7044
	GFJGFKCGLKM                                uint16 = 7046
	KGNHCCPMBLH                                uint16 = 7047
	KMBIPNPCADD                                uint16 = 7048
	JAJOKJHMGBJ                                uint16 = 7049
	OAFHJFCNGKJ                                uint16 = 7050
	MEJLOBJPIJN                                uint16 = 7052
	IJMCPCMPJME                                uint16 = 7054
	CFFJCAECFMD                                uint16 = 7055
	BJOLJJALLLL                                uint16 = 7056
	DJDOLLMENFG                                uint16 = 7057
	KGGEIGMPLNL                                uint16 = 7058
	KAGABFDOIAO                                uint16 = 7059
	JKGPHJNFDIA                                uint16 = 7060
	DPCPMOLHGME                                uint16 = 7061
	KOIEKODNAGI                                uint16 = 7062
	KOEBOMPEPJK                                uint16 = 7064
	PCFGODKODHM                                uint16 = 7065
	AECJKLDGHJK                                uint16 = 7066
	ANHCFIPGLDE                                uint16 = 7067
	OHGHIBHDKLD                                uint16 = 7068
	JCAJLACNKDM                                uint16 = 7069
	BGFOECMPAGB                                uint16 = 7070
	LOEKMDFOIOF                                uint16 = 7071
	PCNKHELGLHJ                                uint16 = 7072
	BAOCKGBCLID                                uint16 = 7073
	OFMKJHFPBNK                                uint16 = 7075
	KKDNFLDMGLM                                uint16 = 7076
	LFKLOAPMFKP                                uint16 = 7077
	BBPBBBDDKFI                                uint16 = 7078
	MNCBHMNLIFG                                uint16 = 7079
	DABEMHCDIEC                                uint16 = 7080
	OBEIGOMKPFB                                uint16 = 7081
	MJPJNDCKOPF                                uint16 = 7083
	ANLCHKLHCIE                                uint16 = 7086
	LFMGJLCBAFB                                uint16 = 7087
	PJGMMJNDHMJ                                uint16 = 7088
	GCPJPEJMDAP                                uint16 = 7089
	PKBAGKLFFGK                                uint16 = 7090
	EDIDBDGJIAF                                uint16 = 7091
	DCKFGKKDKEB                                uint16 = 7092
	FICMEMECAJK                                uint16 = 7093
	GPKACHPPGKI                                uint16 = 7094
	AFPAEBGOLHG                                uint16 = 7095
	GLDPEJDGFEN                                uint16 = 7096
	HHFJJDMGIIP                                uint16 = 7098
	EPLPNAPFBHM                                uint16 = 7099
	DDBDPDMFLME                                uint16 = 7100
	GetEraFlipperDataCsReq                     uint16 = 7102
	ResetEraFlipperDataCsReq                   uint16 = 7105
	PEJJABDLOIH                                uint16 = 7107
	EAOKBOCDHHG                                uint16 = 7109
	KJLCMECMHHE                                uint16 = 7110
	AGOKLPNCBDN                                uint16 = 7112
	BKHONFCFGKB                                uint16 = 7113
	LIHDIHMOLGB                                uint16 = 7114
	HHOHOHHFJLF                                uint16 = 7116
	KMLLPPEIBOP                                uint16 = 7117
	EraFlipperDataChangeScNotify               uint16 = 7119
	MBJFFJDBAJJ                                uint16 = 7120
	GDPEMIMGBNI                                uint16 = 7123
	CLBIFDFPJNK                                uint16 = 7125
	HOMKHAKNBCD                                uint16 = 7126
	DCFMOCENKKJ                                uint16 = 7127
	HMDKLFGHGMG                                uint16 = 7131
	ChangeEraFlipperDataCsReq                  uint16 = 7133
	ResetEraFlipperDataScRsp                   uint16 = 7135
	FKCIBEMNKFE                                uint16 = 7136
	EnterEraFlipperRegionScRsp                 uint16 = 7148
	OLPAAHBAHMJ                                uint16 = 7151
	CNLBFFIELGM                                uint16 = 7152
	ADEBBLIHHNO                                uint16 = 7153
	FPCCJFBKNFI                                uint16 = 7157
	FDECMDAKEDC                                uint16 = 7160
	DOLBDKNMMEL                                uint16 = 7161
	GPHFMAKAMPG                                uint16 = 7162
	GPBIAGAOKKB                                uint16 = 7163
	KCNPMHJGHNM                                uint16 = 7167
	FJMDLCBMMFJ                                uint16 = 7170
	ClockParkGetInfoScRsp                      uint16 = 7202
	BPMLEFKKJCG                                uint16 = 7207
	FEIPJNNPMOC                                uint16 = 7209
	JLMEMDMFKGM                                uint16 = 7210
	BMDPJKOGJDM                                uint16 = 7212
	NNKEMLDCGGP                                uint16 = 7213
	LFODGNFBPOB                                uint16 = 7217
	ClockParkGetInfoCsReq                      uint16 = 7219
	FKNHMACBHLP                                uint16 = 7225
	CCHIMMEMKBO                                uint16 = 7231
	JIOBHLOPKHB                                uint16 = 7233
	FAGPOLDODGP                                uint16 = 7235
	HCBIODPEOHE                                uint16 = 7240
	EBGILDGNGCL                                uint16 = 7243
	JMCMMNKANEO                                uint16 = 7246
	HPDKMOJIGNE                                uint16 = 7250
	DOEAMNCKPNN                                uint16 = 7252
	CGAAGOKHIEM                                uint16 = 7255
	KFAGMGMMHGC                                uint16 = 7269
	CJDINMHAICJ                                uint16 = 7275
	GKIALGMGJND                                uint16 = 7281
	GBFLDCLCEDO                                uint16 = 7283
	LMAOLBGBFHH                                uint16 = 7285
	HCKPAANCJEH                                uint16 = 7298
	StartMatchScRsp                            uint16 = 7302
	CancelMatchScRsp                           uint16 = 7305
	GetCrossInfoScRsp                          uint16 = 7331
	GetCrossInfoCsReq                          uint16 = 7333
	MatchResultScNotify                        uint16 = 7335
	StartMatchCsReq                            uint16 = 7348
	EMIOJOLOCPM                                uint16 = 7352
	IGKLAOPJFOG                                uint16 = 7355
	GEPPCEOCLOG                                uint16 = 7357
	PIHIECEMCCB                                uint16 = 7359
	DMHLEHMCMJA                                uint16 = 7360
	HIIENOHBGPG                                uint16 = 7362
	HDKFBPGAOAM                                uint16 = 7363
	OCHFADEBKHA                                uint16 = 7367
	ELOCPGINHJC                                uint16 = 7369
	FBIALGFBAKH                                uint16 = 7370
	IFAPOJIDINH                                uint16 = 7375
	DBLHNDBNJHF                                uint16 = 7377
	EOGBKAIKCKE                                uint16 = 7381
	NNIEBHKIIIP                                uint16 = 7383
	GOGOMFCOHJA                                uint16 = 7384
	HINLJEPMAEO                                uint16 = 7385
	EKHNLDEELLH                                uint16 = 7390
	KDHMNKGIADL                                uint16 = 7393
	AENCOEBKFMA                                uint16 = 7396
	NMIOCHECKLE                                uint16 = 7398
	DAADKJADMPP                                uint16 = 7400
	NDNODNGCIFP                                uint16 = 7402
	ANFGNJFOFFI                                uint16 = 7405
	JFHBMCMGGPG                                uint16 = 7407
	BEHGLJAPOJF                                uint16 = 7409
	HEOOEKCDMBA                                uint16 = 7412
	KFPNNLAGCON                                uint16 = 7413
	OLIPBINKPHL                                uint16 = 7417
	IPBNIMPLEBD                                uint16 = 7419
	GOFPCLNNHLG                                uint16 = 7420
	AGLLFHMLCFF                                uint16 = 7425
	BNPJMLIIJHE                                uint16 = 7427
	EMKJJGADDLD                                uint16 = 7431
	EEMDHJFINBD                                uint16 = 7433
	LNBENIECPOI                                uint16 = 7434
	BODCPLJDPGF                                uint16 = 7435
	PENCGFNCOEI                                uint16 = 7443
	BBGHLIMPOBH                                uint16 = 7446
	LCNHIJFMBAI                                uint16 = 7448
	NJKMOLIKBAK                                uint16 = 7451
	CHFOOHMJNPO                                uint16 = 7452
	HNMPDDELBMP                                uint16 = 7453
	CKIAKLCJJLH                                uint16 = 7454
	JNPHKPDLOIG                                uint16 = 7455
	GGLAHPHLDCO                                uint16 = 7456
	NNEDJEHPIIE                                uint16 = 7457
	AJBANMKILCG                                uint16 = 7458
	EGCMACJOKME                                uint16 = 7460
	EMALGPDKFAD                                uint16 = 7469
	JALNCLCCNJG                                uint16 = 7470
	LOKDBDKFNJF                                uint16 = 7471
	JIFBPGHCAFF                                uint16 = 7473
	KPLBPAOOCJI                                uint16 = 7474
	ANMLGADENCM                                uint16 = 7475
	MHJDEDOPLLJ                                uint16 = 7476
	JKMPIABFGIO                                uint16 = 7477
	BFGJOPMMELH                                uint16 = 7478
	EMFHGILAAMN                                uint16 = 7479
	CMPCDMFEEBH                                uint16 = 7480
	IMFEIOIEJMP                                uint16 = 7481
	PCGEBNILLME                                uint16 = 7482
	JOMHCPMBFBB                                uint16 = 7483
	IPONIEACFBD                                uint16 = 7484
	CKLNHPPPAHM                                uint16 = 7485
	CPJDLAKBGAF                                uint16 = 7486
	DFCHACFIOCP                                uint16 = 7487
	EGPNAGDIMBK                                uint16 = 7490
	ELACOJOEBMC                                uint16 = 7491
	JLADNFEAAEI                                uint16 = 7493
	PABCOMDNKOG                                uint16 = 7494
	MBMAOKAKEPA                                uint16 = 7495
	MEGEDCOBMKC                                uint16 = 7496
	BCAFEPLNPGG                                uint16 = 7497
	KPIKJKNMFKE                                uint16 = 7498
	PGGINNBANBD                                uint16 = 7500
	ContentPackageGetDataScRsp                 uint16 = 7502
	ContentPackageGetDataCsReq                 uint16 = 7519
	ContentPackageTransferScNotify             uint16 = 7533
	ContentPackageSyncDataScNotify             uint16 = 7548
	MIHHADFCMLC                                uint16 = 7551
	FDNIEBAOINP                                uint16 = 7552
	DFDDNHAMHHC                                uint16 = 7553
	IDIEFCGCKBD                                uint16 = 7555
	AHAMDHHLPEN                                uint16 = 7556
	FPDOEGNIAAK                                uint16 = 7557
	AFBBLGEHOBE                                uint16 = 7559
	OKHPABFGLOG                                uint16 = 7560
	GHEJNPKMHCP                                uint16 = 7561
	GEKLGPLOBOE                                uint16 = 7562
	BPIPPBJBFBE                                uint16 = 7563
	NFKANNMMPEH                                uint16 = 7567
	ICIBMIJKHAK                                uint16 = 7570
	MusicRhythmStartLevelCsReq                 uint16 = 7575
	MusicRhythmFinishLevelCsReq                uint16 = 7577
	MusicRhythmMaxDifficultyLevelsUnlockNotify uint16 = 7579
	MusicRhythmDataScRsp                       uint16 = 7580
	MusicRhythmUnlockSongNotify                uint16 = 7582
	MusicRhythmStartLevelScRsp                 uint16 = 7583
	MusicRhythmFinishLevelScRsp                uint16 = 7587
	MusicRhythmSaveSongConfigDataScRsp         uint16 = 7588
	MusicRhythmUnlockTrackScNotify             uint16 = 7590
	MusicRhythmDataCsReq                       uint16 = 7593
	MusicRhythmUnlockSongSfxScNotify           uint16 = 7594
	MusicRhythmSaveSongConfigDataCsReq         uint16 = 7597
	CurPetChangedScNotify                      uint16 = 7605
	GetPetDataScRsp                            uint16 = 7607
	SummonPetScRsp                             uint16 = 7609
	GetPetDataCsReq                            uint16 = 7617
	RecallPetCsReq                             uint16 = 7620
	RecallPetScRsp                             uint16 = 7622
	SummonPetCsReq                             uint16 = 7624
	WorldUnlockScRsp                           uint16 = 7626
	WorldUnlockCsReq                           uint16 = 7627
	LCAMCNOFPMP                                uint16 = 7652
	LMAHEGLKOJJ                                uint16 = 7655
	PKKLLLANPFG                                uint16 = 7669
	PGFGEEADLJP                                uint16 = 7675
	FKDEBAICAAO                                uint16 = 7681
	PBNFEGELHPB                                uint16 = 7683
	BAEODIMMFGE                                uint16 = 7685
	KGCHIHCEINK                                uint16 = 7698
	CLOAAFAPKPG                                uint16 = 7703
	IGHONOMDLHL                                uint16 = 7706
	PLPOBHLIKLC                                uint16 = 7709
	LLAAGPNGDAI                                uint16 = 7710
	PBCFOGEGLMB                                uint16 = 7711
	BBNOJFFAPCP                                uint16 = 7713
	MNEBHGMAGGD                                uint16 = 7714
	APEGFHBBGDD                                uint16 = 7715
	DMPAINJEGCD                                uint16 = 7717
	DFAKLHGPBCJ                                uint16 = 7719
	DMKHEJINBHI                                uint16 = 7720
	PAOFAEHLGEJ                                uint16 = 7726
	ALMEAJGADBP                                uint16 = 7730
	JDOFCFELEMH                                uint16 = 7733
	LHGOAKJKPDF                                uint16 = 7735
	FBKGCCFALGG                                uint16 = 7737
	GOILEIMINDK                                uint16 = 7738
	NKPDDCNPJPE                                uint16 = 7745
	BAFECPJFFJJ                                uint16 = 7747
	OFMJLFKGBEA                                uint16 = 7749
	GLOPKOBPOIB                                uint16 = 7751
	LAMFNNNCNNL                                uint16 = 7752
	AMEJIHNHFID                                uint16 = 7755
	EKGPBLKOCFH                                uint16 = 7757
	OKHKBMNLLIL                                uint16 = 7761
	FBAIJOOPECG                                uint16 = 7762
	OKFNCCGBHHI                                uint16 = 7766
	AMNPFHHPMGB                                uint16 = 7767
	DPOEEJJBLEC                                uint16 = 7769
	POFMBILCPMK                                uint16 = 7773
	FKEOCMNDHAN                                uint16 = 7776
	CFGHHOMOFJF                                uint16 = 7777
	GEDKMFICABA                                uint16 = 7788
	PNJNIGHFNDA                                uint16 = 7789
	FFBNNEMNAIG                                uint16 = 7791
	GIFKNBFOPDP                                uint16 = 7792
	FOGJMOKNDCL                                uint16 = 7795
	CNNONAOENKL                                uint16 = 7798
	AKNPOFCGAEA                                uint16 = 8001
	TrainPartyGetDataScRsp                     uint16 = 8009
	TrainPartyEnterCsReq                       uint16 = 8010
	TrainPartyGamePlaySettleNotify             uint16 = 8015
	TrainPartyUseCardScRsp                     uint16 = 8017
	MFFCOHOKBHD                                uint16 = 8019
	HANHHJNHLEM                                uint16 = 8026
	BPAHALHPHJB                                uint16 = 8030
	TrainPartyLeaveScRsp                       uint16 = 8035
	TrainPartyUseCardCsReq                     uint16 = 8038
	TrainPartyHandlePendingActionScRsp         uint16 = 8045
	TrainPartyGetDataCsReq                     uint16 = 8047
	IFPJKDDMGEP                                uint16 = 8049
	TrainPartyBuildDiyCsReq                    uint16 = 8050
	TrainPartyMoveScNotify                     uint16 = 8051
	TrainPartyBuildStartStepScRsp              uint16 = 8052
	MOCCJFOHAAB                                uint16 = 8055
	LKOGDOIABBE                                uint16 = 8057
	TrainPartyHandlePendingActionCsReq         uint16 = 8061
	TrainPartyLeaveCsReq                       uint16 = 8062
	TrainPartyEnterScRsp                       uint16 = 8067
	TrainPartyBuildRoomScNotify                uint16 = 8069
	TrainPartyBuildStartStepCsReq              uint16 = 8076
	MJNLNBFAOGA                                uint16 = 8077
	TrainPartyBuildDiyScRsp                    uint16 = 8081
	TrainPartySettleNotify                     uint16 = 8089
	TrainPartyBuildingUpdateNotify             uint16 = 8093
	TrainPartySyncUpdateScNotify               uint16 = 8095
	EGNAJNKNGGG                                uint16 = 8101
	HECMLPGBPEP                                uint16 = 8103
	EJPMLFBNIJL                                uint16 = 8104
	NAJNNIJNAHO                                uint16 = 8105
	SwitchHandDataScRsp                        uint16 = 8106
	JDNLHLGKABA                                uint16 = 8107
	GDLIDFPKEFN                                uint16 = 8109
	APCOIFHMLGJ                                uint16 = 8110
	AGKHKDAEBLL                                uint16 = 8112
	SwitchHandDataCsReq                        uint16 = 8113
	GCPCNFLCLJJ                                uint16 = 8114
	DEKBEKIDNKL                                uint16 = 8115
	MDBHIKBGFAE                                uint16 = 8117
	DAKJIBFCIKD                                uint16 = 8120
	SelectPamSkinScRsp                         uint16 = 8125
	GetPamSkinDataScRsp                        uint16 = 8126
	SelectPamSkinCsReq                         uint16 = 8129
	GetPamSkinDataCsReq                        uint16 = 8133
	UnlockPamSkinScNotify                      uint16 = 8134
	IKALGCPBNDP                                uint16 = 8141
	LODHDFDJIOB                                uint16 = 8142
	ABKDDBIDCJE                                uint16 = 8143
	NMNNAGFIPEP                                uint16 = 8144
	GJOCPAMNIIH                                uint16 = 8145
	JMKEBFCMIPO                                uint16 = 8146
	KINPGIFGGED                                uint16 = 8147
	KMENEDCHAAM                                uint16 = 8149
	HHKCFNGDHKL                                uint16 = 8150
	EAAPHCNFLEK                                uint16 = 8152
	LMHFDBJPCAP                                uint16 = 8153
	BMKBLIHIKIG                                uint16 = 8154
	HKMIMDPJAJN                                uint16 = 8155
	DFELMHMHDOP                                uint16 = 8157
	NKLPPGBOOJN                                uint16 = 8159
	FPMPFCJIOCP                                uint16 = 8160
	JIMMKLIHONA                                uint16 = 8161
	MMPAMLGKGFG                                uint16 = 8162
	ICLGFIHJIAF                                uint16 = 8163
	EABDAHBDKPE                                uint16 = 8164
	HBEDPFHCAPJ                                uint16 = 8165
	BOJIMHGFGEE                                uint16 = 8166
	EBNIKJNJMBD                                uint16 = 8170
	IMEIPJBFLPL                                uint16 = 8173
	OKOMOBEDNPF                                uint16 = 8175
	IOMKBPMHPAM                                uint16 = 8177
	HKFIBMBCPNK                                uint16 = 8180
	UpdateMarkChestScRsp                       uint16 = 8185
	GetMarkChestScRsp                          uint16 = 8186
	UpdateMarkChestCsReq                       uint16 = 8189
	GetMarkChestCsReq                          uint16 = 8193
	MarkChestChangedScNotify                   uint16 = 8194
	EHLDJKJGNCN                                uint16 = 8201
	CHOIPPBOANF                                uint16 = 8202
	OGHKDPJINAK                                uint16 = 8203
	HOJLKBMHNPH                                uint16 = 8204
	DNDEJBMDCMJ                                uint16 = 8205
	DJLLAEPEBHE                                uint16 = 8206
	MPKEEIMNGHI                                uint16 = 8207
	DJBMCJGOLBC                                uint16 = 8208
	FMLCDNJJLMC                                uint16 = 8209
	JBHLHNKJKOF                                uint16 = 8210
	KFDEIFMMHIC                                uint16 = 8211
	FINKJPOOIKG                                uint16 = 8212
	MJOINJMAHBI                                uint16 = 8213
	HFDHGMDBEJF                                uint16 = 8214
	FDBDPJJPDHO                                uint16 = 8215
	FCECPHKCHGK                                uint16 = 8216
	JFJGLMKBAFD                                uint16 = 8217
	OGKFHINGNLC                                uint16 = 8218
	LGNMIJKJKJC                                uint16 = 8219
	DKEOICABNFP                                uint16 = 8220
	HFAJLDKAMEO                                uint16 = 8221
	NDCAPHOBAKC                                uint16 = 8222
	BNOILJLCIOE                                uint16 = 8223
	MIFINMNPMBP                                uint16 = 8224
	CIFAHBDLDED                                uint16 = 8225
	KKMKKBDBMNB                                uint16 = 8226
	AAIANPMCIBG                                uint16 = 8227
	HMKNNAJKAGE                                uint16 = 8228
	DKIJNKLPBFF                                uint16 = 8229
	FOOLDOEIFOC                                uint16 = 8230
	DANCJEECPMN                                uint16 = 8231
	LPAMFAMAFEP                                uint16 = 8232
	BCPCENAKCJK                                uint16 = 8233
	PFHHIGHEDJK                                uint16 = 8234
	AAJFBJHMAON                                uint16 = 8235
	LNOLPAEGLGI                                uint16 = 8236
	HDMEHMJLNAB                                uint16 = 8237
	HEBMHDINMIP                                uint16 = 8238
	LKLAMDJOPCE                                uint16 = 8239
	DOCFADPCGIO                                uint16 = 8240
	OKPIJHJGAFJ                                uint16 = 8241
	KLILCFENABP                                uint16 = 8242
	KIHDMHNCAAL                                uint16 = 8243
	DGKOFIJEHLH                                uint16 = 8244
	OLBOENJMCNL                                uint16 = 8245
	JAGIJPGIBHF                                uint16 = 8246
	HEEDAPEOIHG                                uint16 = 8247
	POKLKNJFJJP                                uint16 = 8248
	OLHNBGKLKDL                                uint16 = 8249
	PLBFHOPECGI                                uint16 = 8250
	RelicSmartWearUpdatePinRelicCsReq          uint16 = 8252
	RelicSmartWearDeletePlanCsReq              uint16 = 8253
	RelicSmartWearUpdatePlanScRsp              uint16 = 8254
	RelicSmartWearAddPlanScRsp                 uint16 = 8255
	RelicSmartWearGetPlanScRsp                 uint16 = 8256
	RelicSmartWearUpdatePinRelicScNotify       uint16 = 8257
	RelicSmartWearAddPlanCsReq                 uint16 = 8259
	RelicSmartWearGetPinRelicCsReq             uint16 = 8260
	RelicSmartWearGetPlanCsReq                 uint16 = 8263
	RelicSmartWearUpdatePlanCsReq              uint16 = 8264
	RelicSmartWearDeletePinRelicScRsp          uint16 = 8266
	RelicSmartWearDeletePlanScRsp              uint16 = 8267
	RelicSmartWearUpdatePinRelicScRsp          uint16 = 8268
	RelicSmartWearGetPinRelicScRsp             uint16 = 8270
	OLJEAPHMGIE                                uint16 = 8271
	GBHLODAEAHH                                uint16 = 8273
	LLBMDLJHJLO                                uint16 = 8274
	BPNJPCNEKBD                                uint16 = 8275
	EGHOBCIMPNH                                uint16 = 8276
	MFANAADHHPG                                uint16 = 8279
	PNIEHBLNHOB                                uint16 = 8282
	KGEIJDNMEKO                                uint16 = 8283
	KCGJFHILEOI                                uint16 = 8284
	OKEKJDFOIKD                                uint16 = 8287
	IAGEJNPIHJG                                uint16 = 8292
	CAOGKPDKPID                                uint16 = 8295
	BDFNNPAPBOJ                                uint16 = 8297
	GEGJIOCOOBI                                uint16 = 8299
	FJBFEBJNBHK                                uint16 = 8302
	EJOBJPJIEMF                                uint16 = 8303
	FIAMCGLOJDE                                uint16 = 8307
	EOPBCNGDODJ                                uint16 = 8309
	FNDNKAACBHD                                uint16 = 8310
	HCBGBMKMAAA                                uint16 = 8315
	BJOCIHICCCM                                uint16 = 8317
	IPJEPOGIDFL                                uint16 = 8321
	CLIPNNJBNFI                                uint16 = 8323
	HMAFHKAEKFG                                uint16 = 8324
	JLHCAEHNBEJ                                uint16 = 8325
	EBKEJLJCMOP                                uint16 = 8333
	KNCIJNEPLHN                                uint16 = 8336
	JCBDEKIGMGC                                uint16 = 8338
	SyncRechargeBenefitInfoScNotify            uint16 = 8363
	GetRechargeBenefitInfoScRsp                uint16 = 8364
	TakeRechargeGiftRewardScRsp                uint16 = 8365
	GetRechargeGiftInfoScRsp                   uint16 = 8366
	TakeRechargeGiftRewardCsReq                uint16 = 8369
	TakeRechargeBenefitRewardScRsp             uint16 = 8372
	GetRechargeGiftInfoCsReq                   uint16 = 8373
	GetRechargeBenefitInfoCsReq                uint16 = 8374
	TakeRechargeBenefitRewardCsReq             uint16 = 8377
	LLIEMGCGHDG                                uint16 = 8383
	LPIKOGPOMNI                                uint16 = 8384
	PPMGKFEDNOL                                uint16 = 8385
	GNLLJMIMNPC                                uint16 = 8386
	FPMFIIPPIMO                                uint16 = 8389
	HDBILLENKMG                                uint16 = 8393
	KIFFLDHFMLL                                uint16 = 8394
	BHHDIGKNGGK                                uint16 = 8397
	GHBGNOKPDKL                                uint16 = 8401
	IKOMIMHGIIN                                uint16 = 8404
	JCMOOEBEONF                                uint16 = 8405
	FEMLPCAFCFE                                uint16 = 8406
	KPBGJHIHKHA                                uint16 = 8410
	EPMNHLAEJOD                                uint16 = 8413
	FIGGHKDLDEH                                uint16 = 8415
	KHNMPOADPFM                                uint16 = 8416
	DBNEGFOMOJG                                uint16 = 8424
	BEOPBDOOPPD                                uint16 = 8427
	BHFCGFIIPEK                                uint16 = 8430
	IEHCLHCBIAP                                uint16 = 8431
	BKLFEDBGIOP                                uint16 = 8432
	KMEOJGKPNIC                                uint16 = 8433
	GPIBEDECNDP                                uint16 = 8435
	IJNCIDCCNFG                                uint16 = 8438
	CJHKHIJCBDL                                uint16 = 8440
	KPMLKCDIFBE                                uint16 = 8441
	JOCPDPNOBNA                                uint16 = 8442
	DLKHMMLJHCA                                uint16 = 8443
	FJDAGBIJMEP                                uint16 = 8445
	EKOOANIDBPP                                uint16 = 8446
	PGHHJGLLOKP                                uint16 = 8448
	FCMPHKJEBJA                                uint16 = 8453
	NCGOHBCBFMH                                uint16 = 8456
	EHBEPEFDPHC                                uint16 = 8459
	PBPIGABCJED                                uint16 = 8465
	EKELNGFIDJD                                uint16 = 8466
	NFPBIDKBCLM                                uint16 = 8468
	CNOINGLLHMB                                uint16 = 8472
	MMJAGLCLIPO                                uint16 = 8474
	GOIIALEIBOB                                uint16 = 8479
	HEHOBIPOHCP                                uint16 = 8480
	MIFGNJIMLKO                                uint16 = 8485
	GLFKBPCDNLM                                uint16 = 8486
	AHBFFEPLBMD                                uint16 = 8487
	CBPGKBBBEEP                                uint16 = 8496
	HHGHILJDNAI                                uint16 = 8497
	GDBEHBFBKKD                                uint16 = 8500
	MMJBHOMPNNH                                uint16 = 8505
	NDLGFKKELLN                                uint16 = 8506
	KNGCCCKIGDF                                uint16 = 8507
	GFDGPBELEGB                                uint16 = 8508
	DLKCLHJNEGB                                uint16 = 8511
	KFHNBIKNLBH                                uint16 = 8515
	KMKOFIKGLJD                                uint16 = 8516
	OKJBBHPFFAN                                uint16 = 8518
	HJFBJFNPDKB                                uint16 = 8519
	JDJFNMFEHNG                                uint16 = 8523
	ABJBJOCBPLH                                uint16 = 8527
	CGDADAACBFO                                uint16 = 8528
	FCDPNMDNGFC                                uint16 = 8529
	NENOCOGHFON                                uint16 = 8531
	DMKLPOKPIGE                                uint16 = 8532
	FENMMJBPKKI                                uint16 = 8533
	FAGAFKIDIKB                                uint16 = 8536
	LEDECKHLPJN                                uint16 = 8538
	GMCAIKPAGMC                                uint16 = 8541
	GKILPAEELJH                                uint16 = 8548
	KACCPALJOFK                                uint16 = 8549
	NDAJHDMMCOI                                uint16 = 8550
	IBOBBDILLEF                                uint16 = 8555
	NBKMJNMLBBO                                uint16 = 8557
	NPFMIPDCLOI                                uint16 = 8559
	HEHLILCFMME                                uint16 = 8563
	GMAAJCPDFBF                                uint16 = 8566
	BNANEKNNMGD                                uint16 = 8567
	DELMBJCMFAP                                uint16 = 8568
	LCNHOJPJKHM                                uint16 = 8571
	PNNJIKHKIHJ                                uint16 = 8575
	MGCCIGAGECD                                uint16 = 8576
	DOEKJBOHGGD                                uint16 = 8577
	NOFKLGLKCAL                                uint16 = 8580
	BANKFKMLPKP                                uint16 = 8586
	OFHPHMDAEAB                                uint16 = 8588
	PGKDCDONKEI                                uint16 = 8589
	NDHPMCECJDN                                uint16 = 8595
	EBOAHJGEKMP                                uint16 = 8597
	ONBJLECGIID                                uint16 = 8598
	GetChallengePeakDataScRsp                  uint16 = 8902
	StartChallengePeakScRsp                    uint16 = 8905
	GetCurChallengePeakScRsp                   uint16 = 8909
	SetChallengePeakMobLineupAvatarCsReq       uint16 = 8910
	GetCurChallengePeakCsReq                   uint16 = 8912
	SetChallengePeakBossHardModeCsReq          uint16 = 8913
	ConfirmChallengePeakSettleScRsp            uint16 = 8916
	TakeChallengePeakRewardScRsp               uint16 = 8917
	GetChallengePeakDataCsReq                  uint16 = 8919
	ReStartChallengePeakScRsp                  uint16 = 8920
	ConfirmChallengePeakSettleCsReq            uint16 = 8923
	ChallengePeakGroupDataUpdateScNotify       uint16 = 8925
	ChallengePeakSettleScNotify                uint16 = 8931
	LeaveChallengePeakScRsp                    uint16 = 8933
	ReStartChallengePeakCsReq                  uint16 = 8934
	LeaveChallengePeakCsReq                    uint16 = 8935
	SetChallengePeakBossHardModeScRsp          uint16 = 8943
	TakeChallengePeakRewardCsReq               uint16 = 8946
	StartChallengePeakCsReq                    uint16 = 8948
	SetChallengePeakMobLineupAvatarScRsp       uint16 = 8950
	FKBJJIMKFOK                                uint16 = 8951
	NPGONPEDINB                                uint16 = 8954
	IDIKMHHOIIC                                uint16 = 8955
	NJHPBNBHOLL                                uint16 = 8956
	CAICLJOILEK                                uint16 = 8959
	GJMDFCALIOL                                uint16 = 8960
	NDINECKDMDO                                uint16 = 8962
	EKNCLPMCCNE                                uint16 = 8963
	JNKIOGAEJDG                                uint16 = 8964
	FPLOMPFOFPL                                uint16 = 8965
	JIAMLEMCLHD                                uint16 = 8969
	GOHAIMEMEFG                                uint16 = 8970
	MLGECIDINHM                                uint16 = 9002
	LKJMMAJJCBK                                uint16 = 9005
	IMKPBFNLMBJ                                uint16 = 9007
	MCHDFOHHBLF                                uint16 = 9009
	PKLFFNMIIAN                                uint16 = 9010
	JBFCBGJADOA                                uint16 = 9012
	IJJLBEKHICO                                uint16 = 9013
	OHFABPDBIDE                                uint16 = 9017
	PFCGHODONEK                                uint16 = 9019
	FBPHNLFDFAE                                uint16 = 9025
	HJJAFGIBNDI                                uint16 = 9027
	GIGAMCJGJKP                                uint16 = 9031
	AJDBHHLKLBB                                uint16 = 9033
	GMJIHGDONDE                                uint16 = 9035
	MHOAJAJNBGH                                uint16 = 9043
	JBGBOMDDOEN                                uint16 = 9046
	JIHBNKNCJKH                                uint16 = 9048
	ANPNFCFDDLA                                uint16 = 9050
	EnterElationActivityStageScRsp             uint16 = 9101
	GetElationActivityDataCsReq                uint16 = 9102
	ElationActivityBattleEndScNotify           uint16 = 9103
	GetElationActivityDataScRsp                uint16 = 9107
	EnterElationActivityStageCsReq             uint16 = 9110
	GetActivityRewardCountDataScRsp            uint16 = 9111
	GetActivityHotDataCsReq                    uint16 = 9112
	GetActivityHotDataScRsp                    uint16 = 9117
	GetActivityRewardCountDataCsReq            uint16 = 9120
	FHPKLLOPKOC                                uint16 = 9121
	JEFONNOLAEB                                uint16 = 9122
	KEMCCEIJAKD                                uint16 = 9123
	EHGALIKLNHA                                uint16 = 9124
	BOFANBLCKGI                                uint16 = 9125
	HGCHMEOGOMA                                uint16 = 9126
	FCPEMJIJNPJ                                uint16 = 9127
	MENBJHDINMP                                uint16 = 9129
	FFMDKKCIHMA                                uint16 = 9130
	NBJPOFHICID                                uint16 = 9132
	NDLOBADOAAO                                uint16 = 9133
	LEIBJEKBPBJ                                uint16 = 9134
	HMJICDOPOLH                                uint16 = 9135
	KDCGLMGICMG                                uint16 = 9137
	GKMKCEAMKPF                                uint16 = 9138
	OKPOACNHJEK                                uint16 = 9142
	AMFBCEOPCDP                                uint16 = 9143
	PABIPMPHMHC                                uint16 = 9144
	CCIENKDNOAE                                uint16 = 9146
	AGJLOBMFKNC                                uint16 = 9149
	DMGIOHDPGIO                                uint16 = 9150
	BNINEBOOBCA                                uint16 = 9211
	DAEDKDPBMIM                                uint16 = 9212
	KAMEILLBJOF                                uint16 = 9214
	MIIEPDAIGCH                                uint16 = 9217
	FPEONCPAMOI                                uint16 = 9219
	OIHHMIBGPIA                                uint16 = 9220
	GBLKBGALMIB                                uint16 = 9221
	PBGEKEEIHCI                                uint16 = 9223
	AEKEJPOHHLO                                uint16 = 9224
	HBPLLJKPEIA                                uint16 = 9226
	NENFNAACLGL                                uint16 = 9227
	MIHENFCLPHE                                uint16 = 9228
	CJGIAPJKIDL                                uint16 = 9229
	AFNHGDHFPGO                                uint16 = 9230
	NACKBGDHDIC                                uint16 = 9233
	HGDLDLDKFJD                                uint16 = 9234
	IMNEGOKAGHM                                uint16 = 9235
	GDNGCFMFJIA                                uint16 = 9236
	NIFBFFOBEAN                                uint16 = 9237
	DNKCJELFHOK                                uint16 = 9239
	OMMKLGOJJMB                                uint16 = 9240
	EHPCNCMNBOH                                uint16 = 9241
	GMDGFJLKGKF                                uint16 = 9243
	GCAJBFICCCL                                uint16 = 9244
	HAPPJBCBMGC                                uint16 = 9245
	IAGEPPHBBPC                                uint16 = 9246
	NODCKLPHMIM                                uint16 = 9249
	IKIBCLKMNJJ                                uint16 = 9250
	MJIMHKNAMCD                                uint16 = 9253
	DNLBFEELOOA                                uint16 = 9254
	GNDNNGOMNPB                                uint16 = 9255
	GGKNJFGLKKE                                uint16 = 9256
	OCKFHJAHCCJ                                uint16 = 9257
	AIPGPJMNOCB                                uint16 = 9260
	FightEnterScRsp                            uint16 = 30009
	FightHeartBeatScRsp                        uint16 = 30011
	FightKickOutScNotify                       uint16 = 30017
	FightLeaveScNotify                         uint16 = 30038
	FightSessionStopScNotify                   uint16 = 30089
	FightGeneralScNotify                       uint16 = 30091
	FightGeneralScRsp                          uint16 = 30095
	FNDJDJNEAEA                                uint16 = 30109
	OEBPDLGBCBB                                uint16 = 30110
	LIGDMICINAB                                uint16 = 30114
	IGEIIOKAKLG                                uint16 = 30117
	GBONFNCICAL                                uint16 = 30120
	BLDPPINBDFC                                uint16 = 30126
	FENMJKPFKFO                                uint16 = 30137
	FCJHEBMDCHA                                uint16 = 30138
	GGHDPLPGCMN                                uint16 = 30145
	ELNGDGGHLPL                                uint16 = 30151
	JKIAHGLADJF                                uint16 = 30152
	CCONBOOCEDK                                uint16 = 30161
	CONOIMJDDPB                                uint16 = 30162
	BGEMCKFPJFB                                uint16 = 30167
	ICIELFFCAEG                                uint16 = 30181
	BCICFMMJJOI                                uint16 = 30189
	DFAKFBEFHBI                                uint16 = 30191
	EJJMMHMPGNF                                uint16 = 30192
	BIHJPPJHOHH                                uint16 = 30198
)

var names = map[uint16]string{
	NONE:                                      "NONE",
	SetPlayerInfoCsReq:                        "SetPlayerInfoCsReq",
	UpdatePsnSettingsInfoCsReq:                "UpdatePsnSettingsInfoCsReq",
	GetGameStateServiceConfigScRsp:            "GetGameStateServiceConfigScRsp",
	SetGenderCsReq:                            "SetGenderCsReq",
	GetBasicInfoScRsp:                         "GetBasicInfoScRsp",
	PlayerLoginScRsp:                          "PlayerLoginScRsp",
	AntiAddictScNotify:                        "AntiAddictScNotify",
	PlayerGetTokenScRsp:                       "PlayerGetTokenScRsp",
	SetLanguageScRsp:                          "SetLanguageScRsp",
	GetSecretKeyInfoCsReq:                     "GetSecretKeyInfoCsReq",
	GmTalkScNotify:                            "GmTalkScNotify",
	UpdatePlayerSettingCsReq:                  "UpdatePlayerSettingCsReq",
	ClientObjDownloadDataScNotify:             "ClientObjDownloadDataScNotify",
	UpdatePsnSettingsInfoScRsp:                "UpdatePsnSettingsInfoScRsp",
	GetBasicInfoCsReq:                         "GetBasicInfoCsReq",
	GateServerScNotify:                        "GateServerScNotify",
	FeatureSwitchClosedScNotify:               "FeatureSwitchClosedScNotify",
	SetGameplayBirthdayCsReq:                  "SetGameplayBirthdayCsReq",
	QueryProductInfoCsReq:                     "QueryProductInfoCsReq",
	AceAntiCheaterCsReq:                       "AceAntiCheaterCsReq",
	GetLevelRewardTakenListCsReq:              "GetLevelRewardTakenListCsReq",
	ClientDownloadDataScNotify:                "ClientDownloadDataScNotify",
	RegionStopScNotify:                        "RegionStopScNotify",
	PlayerLogoutCsReq:                         "PlayerLogoutCsReq",
	PlayerHeartBeatCsReq:                      "PlayerHeartBeatCsReq",
	ReserveStaminaExchangeCsReq:               "ReserveStaminaExchangeCsReq",
	PlayerForceSyncGameStateFinishCsReq:       "PlayerForceSyncGameStateFinishCsReq",
	ClientObjUploadScRsp:                      "ClientObjUploadScRsp",
	GetLevelRewardCsReq:                       "GetLevelRewardCsReq",
	PlayerLoginCsReq:                          "PlayerLoginCsReq",
	GetVideoVersionKeyScRsp:                   "GetVideoVersionKeyScRsp",
	ExchangeStaminaScRsp:                      "ExchangeStaminaScRsp",
	PlayerGetTokenCsReq:                       "PlayerGetTokenCsReq",
	ExchangeStaminaCsReq:                      "ExchangeStaminaCsReq",
	PlayerLoginFinishScRsp:                    "PlayerLoginFinishScRsp",
	SetPlayerInfoScRsp:                        "SetPlayerInfoScRsp",
	SetLanguageCsReq:                          "SetLanguageCsReq",
	ReserveStaminaExchangeScRsp:               "ReserveStaminaExchangeScRsp",
	GmTalkScRsp:                               "GmTalkScRsp",
	SetNicknameScRsp:                          "SetNicknameScRsp",
	RetcodeNotify:                             "RetcodeNotify",
	QueryProductInfoScRsp:                     "QueryProductInfoScRsp",
	SetNicknameCsReq:                          "SetNicknameCsReq",
	AceAntiCheaterScRsp:                       "AceAntiCheaterScRsp",
	MonthCardRewardNotify:                     "MonthCardRewardNotify",
	ServerAnnounceNotify:                      "ServerAnnounceNotify",
	UseReserveAndFuelExchangeStaminaScRsp:     "UseReserveAndFuelExchangeStaminaScRsp",
	StaminaInfoScNotify:                       "StaminaInfoScNotify",
	PlayerHeartBeatScRsp:                      "PlayerHeartBeatScRsp",
	GetSecretKeyInfoScRsp:                     "GetSecretKeyInfoScRsp",
	GetAuthkeyCsReq:                           "GetAuthkeyCsReq",
	PlayerLoginFinishCsReq:                    "PlayerLoginFinishCsReq",
	PlayerForceSyncGameStateFinishScRsp:       "PlayerForceSyncGameStateFinishScRsp",
	UpdateFeatureSwitchScNotify:               "UpdateFeatureSwitchScNotify",
	SetRedPointStatusScNotify:                 "SetRedPointStatusScNotify",
	SetGenderScRsp:                            "SetGenderScRsp",
	ClientObjUploadCsReq:                      "ClientObjUploadCsReq",
	GmTalkCsReq:                               "GmTalkCsReq",
	GetLevelRewardScRsp:                       "GetLevelRewardScRsp",
	GetAuthkeyScRsp:                           "GetAuthkeyScRsp",
	PlayerKickOutScNotify:                     "PlayerKickOutScNotify",
	GetVideoVersionKeyCsReq:                   "GetVideoVersionKeyCsReq",
	DailyRefreshNotify:                        "DailyRefreshNotify",
	GetLevelRewardTakenListScRsp:              "GetLevelRewardTakenListScRsp",
	UpdatePlayerSettingScRsp:                  "UpdatePlayerSettingScRsp",
	SetGameplayBirthdayScRsp:                  "SetGameplayBirthdayScRsp",
	PVEBattleResultScRsp:                      "PVEBattleResultScRsp",
	GetCurBattleInfoScRsp:                     "GetCurBattleInfoScRsp",
	QuitBattleScRsp:                           "QuitBattleScRsp",
	SyncClientResVersionScRsp:                 "SyncClientResVersionScRsp",
	QuitBattleCsReq:                           "QuitBattleCsReq",
	ServerSimulateBattleFinishScNotify:        "ServerSimulateBattleFinishScNotify",
	PVEBattleResultCsReq:                      "PVEBattleResultCsReq",
	GetCurBattleInfoCsReq:                     "GetCurBattleInfoCsReq",
	RebattleByClientCsNotify:                  "RebattleByClientCsNotify",
	BattleLogReportScRsp:                      "BattleLogReportScRsp",
	ReBattleAfterBattleLoseCsNotify:           "ReBattleAfterBattleLoseCsNotify",
	SyncClientResVersionCsReq:                 "SyncClientResVersionCsReq",
	QuitBattleScNotify:                        "QuitBattleScNotify",
	GetPreAvatarGrowthInfoCsReq:               "GetPreAvatarGrowthInfoCsReq",
	GetPreAvatarActivityListScRsp:             "GetPreAvatarActivityListScRsp",
	SetAvatarEnhancedIdScRsp:                  "SetAvatarEnhancedIdScRsp",
	SetMultipleAvatarPathsCsReq:               "SetMultipleAvatarPathsCsReq",
	GetAvatarDataScRsp:                        "GetAvatarDataScRsp",
	TakeOffRelicScRsp:                         "TakeOffRelicScRsp",
	UnlockSkilltreeScRsp:                      "UnlockSkilltreeScRsp",
	GrowthTargetAvatarChangedScNotify:         "GrowthTargetAvatarChangedScNotify",
	UnlockAvatarSkinScNotify:                  "UnlockAvatarSkinScNotify",
	AddMultiPathAvatarScNotify:                "AddMultiPathAvatarScNotify",
	AvatarExpUpScRsp:                          "AvatarExpUpScRsp",
	PromoteAvatarScRsp:                        "PromoteAvatarScRsp",
	UnlockAvatarPathScRsp:                     "UnlockAvatarPathScRsp",
	MarkAvatarCsReq:                           "MarkAvatarCsReq",
	SetPlayerOutfitScRsp:                      "SetPlayerOutfitScRsp",
	AvatarSpecialSkilltreeUnlockScNotify:      "AvatarSpecialSkilltreeUnlockScNotify",
	SetAvatarPathCsReq:                        "SetAvatarPathCsReq",
	DressAvatarSkinCsReq:                      "DressAvatarSkinCsReq",
	AvatarPathChangedNotify:                   "AvatarPathChangedNotify",
	TakeOffRelicCsReq:                         "TakeOffRelicCsReq",
	AvatarExpUpCsReq:                          "AvatarExpUpCsReq",
	TakeOffEquipmentScRsp:                     "TakeOffEquipmentScRsp",
	TakeOffAvatarSkinCsReq:                    "TakeOffAvatarSkinCsReq",
	GetAvatarDataCsReq:                        "GetAvatarDataCsReq",
	GetPreAvatarGrowthInfoScRsp:               "GetPreAvatarGrowthInfoScRsp",
	RankUpAvatarScRsp:                         "RankUpAvatarScRsp",
	UnlockSkilltreeCsReq:                      "UnlockSkilltreeCsReq",
	RankUpAvatarCsReq:                         "RankUpAvatarCsReq",
	SetPlayerOutfitCsReq:                      "SetPlayerOutfitCsReq",
	MarkAvatarScRsp:                           "MarkAvatarScRsp",
	TakeOffEquipmentCsReq:                     "TakeOffEquipmentCsReq",
	TakePromotionRewardScRsp:                  "TakePromotionRewardScRsp",
	SetAvatarPathScRsp:                        "SetAvatarPathScRsp",
	TakePromotionRewardCsReq:                  "TakePromotionRewardCsReq",
	SetAvatarEnhancedIdCsReq:                  "SetAvatarEnhancedIdCsReq",
	SetGrowthTargetAvatarScRsp:                "SetGrowthTargetAvatarScRsp",
	SetGrowthTargetAvatarCsReq:                "SetGrowthTargetAvatarCsReq",
	AddAvatarScNotify:                         "AddAvatarScNotify",
	DressRelicAvatarCsReq:                     "DressRelicAvatarCsReq",
	UnlockAvatarPathCsReq:                     "UnlockAvatarPathCsReq",
	GetPreAvatarActivityListCsReq:             "GetPreAvatarActivityListCsReq",
	PromoteAvatarCsReq:                        "PromoteAvatarCsReq",
	DressAvatarScRsp:                          "DressAvatarScRsp",
	TakeOffAvatarSkinScRsp:                    "TakeOffAvatarSkinScRsp",
	DressRelicAvatarScRsp:                     "DressRelicAvatarScRsp",
	DressAvatarCsReq:                          "DressAvatarCsReq",
	SetMultipleAvatarPathsScRsp:               "SetMultipleAvatarPathsScRsp",
	DressAvatarSkinScRsp:                      "DressAvatarSkinScRsp",
	GetWaypointScRsp:                          "GetWaypointScRsp",
	GetChapterScRsp:                           "GetChapterScRsp",
	SetCurWaypointScRsp:                       "SetCurWaypointScRsp",
	BNJCKFMEDCC:                               "BNJCKFMEDCC",
	NFOMIJJIGFA:                               "NFOMIJJIGFA",
	GetMarkItemListCsReq:                      "GetMarkItemListCsReq",
	RelicReforgeConfirmCsReq:                  "RelicReforgeConfirmCsReq",
	SetTurnFoodSwitchScRsp:                    "SetTurnFoodSwitchScRsp",
	SyncTurnFoodNotify:                        "SyncTurnFoodNotify",
	GetRelicFilterPlanCsReq:                   "GetRelicFilterPlanCsReq",
	GetBagScRsp:                               "GetBagScRsp",
	SellItemCsReq:                             "SellItemCsReq",
	LockEquipmentScRsp:                        "LockEquipmentScRsp",
	DestroyItemScRsp:                          "DestroyItemScRsp",
	AddEquipmentScNotify:                      "AddEquipmentScNotify",
	ComposeLimitNumCompleteNotify:             "ComposeLimitNumCompleteNotify",
	PromoteEquipmentScRsp:                     "PromoteEquipmentScRsp",
	RelicReforgeConfirmScRsp:                  "RelicReforgeConfirmScRsp",
	MarkItemScRsp:                             "MarkItemScRsp",
	UseItemScRsp:                              "UseItemScRsp",
	GetRecyleTimeCsReq:                        "GetRecyleTimeCsReq",
	RelicFilterPlanClearNameScNotify:          "RelicFilterPlanClearNameScNotify",
	DiscardRelicCsReq:                         "DiscardRelicCsReq",
	RelicReforgeCsReq:                         "RelicReforgeCsReq",
	AddRelicFilterPlanCsReq:                   "AddRelicFilterPlanCsReq",
	DiscardRelicScRsp:                         "DiscardRelicScRsp",
	ModifyRelicFilterPlanCsReq:                "ModifyRelicFilterPlanCsReq",
	ExchangeHcoinCsReq:                        "ExchangeHcoinCsReq",
	LockRelicScRsp:                            "LockRelicScRsp",
	PromoteEquipmentCsReq:                     "PromoteEquipmentCsReq",
	MarkRelicFilterPlanCsReq:                  "MarkRelicFilterPlanCsReq",
	RelicReforgeScRsp:                         "RelicReforgeScRsp",
	ExpUpEquipmentScRsp:                       "ExpUpEquipmentScRsp",
	ComposeSelectedRelicCsReq:                 "ComposeSelectedRelicCsReq",
	GetBagCsReq:                               "GetBagCsReq",
	GetMarkItemListScRsp:                      "GetMarkItemListScRsp",
	ExpUpRelicCsReq:                           "ExpUpRelicCsReq",
	LockEquipmentCsReq:                        "LockEquipmentCsReq",
	ComposeItemScRsp:                          "ComposeItemScRsp",
	GeneralVirtualItemDataNotify:              "GeneralVirtualItemDataNotify",
	GetRecyleTimeScRsp:                        "GetRecyleTimeScRsp",
	ExpUpEquipmentCsReq:                       "ExpUpEquipmentCsReq",
	RechargeSuccNotify:                        "RechargeSuccNotify",
	DeleteRelicFilterPlanScRsp:                "DeleteRelicFilterPlanScRsp",
	SellItemScRsp:                             "SellItemScRsp",
	ModifyRelicFilterPlanScRsp:                "ModifyRelicFilterPlanScRsp",
	CancelMarkItemNotify:                      "CancelMarkItemNotify",
	DestroyItemCsReq:                          "DestroyItemCsReq",
	DeleteRelicFilterPlanCsReq:                "DeleteRelicFilterPlanCsReq",
	ComposeLimitNumUpdateNotify:               "ComposeLimitNumUpdateNotify",
	ComposeItemCsReq:                          "ComposeItemCsReq",
	MarkItemCsReq:                             "MarkItemCsReq",
	MarkRelicFilterPlanScRsp:                  "MarkRelicFilterPlanScRsp",
	ExpUpRelicScRsp:                           "ExpUpRelicScRsp",
	BatchRankUpEquipmentScRsp:                 "BatchRankUpEquipmentScRsp",
	SetTurnFoodSwitchCsReq:                    "SetTurnFoodSwitchCsReq",
	UseItemCsReq:                              "UseItemCsReq",
	RankUpEquipmentScRsp:                      "RankUpEquipmentScRsp",
	ComposeSelectedRelicScRsp:                 "ComposeSelectedRelicScRsp",
	LockRelicCsReq:                            "LockRelicCsReq",
	RankUpEquipmentCsReq:                      "RankUpEquipmentCsReq",
	GetRelicFilterPlanScRsp:                   "GetRelicFilterPlanScRsp",
	ExchangeHcoinScRsp:                        "ExchangeHcoinScRsp",
	AddRelicFilterPlanScRsp:                   "AddRelicFilterPlanScRsp",
	PlayerSyncScNotify:                        "PlayerSyncScNotify",
	GetStageLineupScRsp:                       "GetStageLineupScRsp",
	SetLineupNameScRsp:                        "SetLineupNameScRsp",
	JoinLineupScRsp:                           "JoinLineupScRsp",
	VirtualLineupTrialAvatarChangeScNotify:    "VirtualLineupTrialAvatarChangeScNotify",
	GetCurLineupDataScRsp:                     "GetCurLineupDataScRsp",
	QuitLineupScRsp:                           "QuitLineupScRsp",
	VirtualLineupDestroyNotify:                "VirtualLineupDestroyNotify",
	SetLineupNameCsReq:                        "SetLineupNameCsReq",
	GetCurLineupDataCsReq:                     "GetCurLineupDataCsReq",
	GetLineupAvatarDataCsReq:                  "GetLineupAvatarDataCsReq",
	ReplaceLineupScRsp:                        "ReplaceLineupScRsp",
	ChangeLineupLeaderScRsp:                   "ChangeLineupLeaderScRsp",
	JoinLineupCsReq:                           "JoinLineupCsReq",
	ChangeLineupLeaderCsReq:                   "ChangeLineupLeaderCsReq",
	SyncLineupNotify:                          "SyncLineupNotify",
	GetAllLineupDataScRsp:                     "GetAllLineupDataScRsp",
	GetAllLineupDataCsReq:                     "GetAllLineupDataCsReq",
	GetLineupAvatarDataScRsp:                  "GetLineupAvatarDataScRsp",
	SwitchLineupIndexCsReq:                    "SwitchLineupIndexCsReq",
	QuitLineupCsReq:                           "QuitLineupCsReq",
	SwapLineupScRsp:                           "SwapLineupScRsp",
	ExtraLineupDestroyNotify:                  "ExtraLineupDestroyNotify",
	SwitchLineupIndexScRsp:                    "SwitchLineupIndexScRsp",
	SwapLineupCsReq:                           "SwapLineupCsReq",
	ReplaceLineupCsReq:                        "ReplaceLineupCsReq",
	GetMailScRsp:                              "GetMailScRsp",
	DelMailScRsp:                              "DelMailScRsp",
	MarkReadMailScRsp:                         "MarkReadMailScRsp",
	TakeMailAttachmentScRsp:                   "TakeMailAttachmentScRsp",
	MarkReadMailCsReq:                         "MarkReadMailCsReq",
	GetMailCsReq:                              "GetMailCsReq",
	DelMailCsReq:                              "DelMailCsReq",
	TakeMailAttachmentCsReq:                   "TakeMailAttachmentCsReq",
	NewMailScNotify:                           "NewMailScNotify",
	GetQuestDataScRsp:                         "GetQuestDataScRsp",
	TakeQuestRewardScRsp:                      "TakeQuestRewardScRsp",
	GetQuestRecordScRsp:                       "GetQuestRecordScRsp",
	TakeQuestRewardCsReq:                      "TakeQuestRewardCsReq",
	TakeQuestOptionalRewardCsReq:              "TakeQuestOptionalRewardCsReq",
	GetQuestDataCsReq:                         "GetQuestDataCsReq",
	FinishQuestScRsp:                          "FinishQuestScRsp",
	TakeQuestOptionalRewardScRsp:              "TakeQuestOptionalRewardScRsp",
	BatchGetQuestDataScRsp:                    "BatchGetQuestDataScRsp",
	GetQuestRecordCsReq:                       "GetQuestRecordCsReq",
	FinishQuestCsReq:                          "FinishQuestCsReq",
	QuestRecordScNotify:                       "QuestRecordScNotify",
	PCLPIMJPADF:                               "PCLPIMJPADF",
	ECHFCICIAGO:                               "ECHFCICIAGO",
	NMFBEMKIBED:                               "NMFBEMKIBED",
	KJEBHAEAPCG:                               "KJEBHAEAPCG",
	DJNJLMODHFB:                               "DJNJLMODHFB",
	GCJNDOAGDIF:                               "GCJNDOAGDIF",
	OKBDGICAGAA:                               "OKBDGICAGAA",
	MFCMNNJGHDJ:                               "MFCMNNJGHDJ",
	NEDMNNMKCFD:                               "NEDMNNMKCFD",
	FOCOPOALNFP:                               "FOCOPOALNFP",
	GPAFHGJOCGH:                               "GPAFHGJOCGH",
	BIAPIHNJHIN:                               "BIAPIHNJHIN",
	DNCIJBOACFO:                               "DNCIJBOACFO",
	HHMDEJFGFKG:                               "HHMDEJFGFKG",
	AILPJKHCGGB:                               "AILPJKHCGGB",
	PJMOBPLICEO:                               "PJMOBPLICEO",
	HFICGFPKIJI:                               "HFICGFPKIJI",
	GetMainMissionCustomValueCsReq:            "GetMainMissionCustomValueCsReq",
	UpdateMainMissionCustomValueCsReq:         "UpdateMainMissionCustomValueCsReq",
	UpdateTrackMainMissionIdScRsp:             "UpdateTrackMainMissionIdScRsp",
	GetMissionDataScRsp:                       "GetMissionDataScRsp",
	SyncTaskCsReq:                             "SyncTaskCsReq",
	AcceptMainMissionScRsp:                    "AcceptMainMissionScRsp",
	SubMissionRewardScNotify:                  "SubMissionRewardScNotify",
	StartFinishSubMissionScNotify:             "StartFinishSubMissionScNotify",
	FinishTalkMissionScRsp:                    "FinishTalkMissionScRsp",
	TeleportToMissionResetPointCsReq:          "TeleportToMissionResetPointCsReq",
	DECPBAPAEGE:                               "DECPBAPAEGE",
	AAEAIIFGMEE:                               "AAEAIIFGMEE",
	FinishTalkMissionCsReq:                    "FinishTalkMissionCsReq",
	MissionGroupWarnScNotify:                  "MissionGroupWarnScNotify",
	GetMissionDataCsReq:                       "GetMissionDataCsReq",
	GetMainMissionCustomValueScRsp:            "GetMainMissionCustomValueScRsp",
	MissionRewardScNotify:                     "MissionRewardScNotify",
	FinishCosumeItemMissionScRsp:              "FinishCosumeItemMissionScRsp",
	UpdateMainMissionCustomValueScRsp:         "UpdateMainMissionCustomValueScRsp",
	TeleportToMissionResetPointScRsp:          "TeleportToMissionResetPointScRsp",
	GetMissionStatusScRsp:                     "GetMissionStatusScRsp",
	GetMissionStatusCsReq:                     "GetMissionStatusCsReq",
	UpdateTrackMainMissionIdCsReq:             "UpdateTrackMainMissionIdCsReq",
	AcceptMainMissionCsReq:                    "AcceptMainMissionCsReq",
	StartFinishMainMissionScNotify:            "StartFinishMainMissionScNotify",
	FinishCosumeItemMissionCsReq:              "FinishCosumeItemMissionCsReq",
	MissionAcceptScNotify:                     "MissionAcceptScNotify",
	FinishedMissionScNotify:                   "FinishedMissionScNotify",
	SyncTaskScRsp:                             "SyncTaskScRsp",
	EnterAdventureScRsp:                       "EnterAdventureScRsp",
	QuickStartCocoonStageScRsp:                "QuickStartCocoonStageScRsp",
	GetFarmStageGachaInfoScRsp:                "GetFarmStageGachaInfoScRsp",
	QuickStartFarmElementScRsp:                "QuickStartFarmElementScRsp",
	GetFarmStageGachaInfoCsReq:                "GetFarmStageGachaInfoCsReq",
	FarmElementSweepScRsp:                     "FarmElementSweepScRsp",
	QuickStartCocoonStageCsReq:                "QuickStartCocoonStageCsReq",
	FarmElementSweepCsReq:                     "FarmElementSweepCsReq",
	QuickStartFarmElementCsReq:                "QuickStartFarmElementCsReq",
	CocoonSweepScRsp:                          "CocoonSweepScRsp",
	CocoonSweepCsReq:                          "CocoonSweepCsReq",
	RecoverAllLineupCsReq:                     "RecoverAllLineupCsReq",
	ReEnterLastElementStageCsReq:              "ReEnterLastElementStageCsReq",
	SetClientPausedCsReq:                      "SetClientPausedCsReq",
	SceneGroupRefreshScNotify:                 "SceneGroupRefreshScNotify",
	UnlockedAreaMapScNotify:                   "UnlockedAreaMapScNotify",
	GetUnlockTeleportCsReq:                    "GetUnlockTeleportCsReq",
	EnteredSceneChangeScNotify:                "EnteredSceneChangeScNotify",
	SceneEntityMoveScRsp:                      "SceneEntityMoveScRsp",
	SpringRefreshScRsp:                        "SpringRefreshScRsp",
	SceneCastSkillScRsp:                       "SceneCastSkillScRsp",
	DeleteSummonUnitScRsp:                     "DeleteSummonUnitScRsp",
	SetCurInteractEntityScRsp:                 "SetCurInteractEntityScRsp",
	SceneReviveAfterRebattleScRsp:             "SceneReviveAfterRebattleScRsp",
	DeleteSummonUnitCsReq:                     "DeleteSummonUnitCsReq",
	InteractPropScRsp:                         "InteractPropScRsp",
	ReEnterLastElementStageScRsp:              "ReEnterLastElementStageScRsp",
	StartCocoonStageCsReq:                     "StartCocoonStageCsReq",
	GetCurSceneInfoScRsp:                      "GetCurSceneInfoScRsp",
	GroupStateChangeCsReq:                     "GroupStateChangeCsReq",
	OpenChestScNotify:                         "OpenChestScNotify",
	GameplayCounterCountDownCsReq:             "GameplayCounterCountDownCsReq",
	GameplayCounterRecoverCsReq:               "GameplayCounterRecoverCsReq",
	EnterSectionCsReq:                         "EnterSectionCsReq",
	SetGroupCustomSaveDataScRsp:               "SetGroupCustomSaveDataScRsp",
	GetSceneMapInfoCsReq:                      "GetSceneMapInfoCsReq",
	UpdateGroupPropertyScRsp:                  "UpdateGroupPropertyScRsp",
	DeactivateFarmElementCsReq:                "DeactivateFarmElementCsReq",
	ReturnLastTownScRsp:                       "ReturnLastTownScRsp",
	ActivateFarmElementCsReq:                  "ActivateFarmElementCsReq",
	SpringRefreshCsReq:                        "SpringRefreshCsReq",
	InteractPropCsReq:                         "InteractPropCsReq",
	UpdateMechanismBarScNotify:                "UpdateMechanismBarScNotify",
	SyncServerSceneChangeNotify:               "SyncServerSceneChangeNotify",
	UnlockTeleportNotify:                      "UnlockTeleportNotify",
	TrainWorldIdChangeScNotify:                "TrainWorldIdChangeScNotify",
	GameplayCounterUpdateScNotify:             "GameplayCounterUpdateScNotify",
	SceneUpdatePositionVersionNotify:          "SceneUpdatePositionVersionNotify",
	SceneEnterStageScRsp:                      "SceneEnterStageScRsp",
	SceneEntityMoveCsReq:                      "SceneEntityMoveCsReq",
	EnterSceneByServerScNotify:                "EnterSceneByServerScNotify",
	RecoverAllLineupScRsp:                     "RecoverAllLineupScRsp",
	SceneCastSkillCostMpScRsp:                 "SceneCastSkillCostMpScRsp",
	SceneCastSkillCsReq:                       "SceneCastSkillCsReq",
	SceneCastSkillCostMpCsReq:                 "SceneCastSkillCostMpCsReq",
	EnterSceneCsReq:                           "EnterSceneCsReq",
	GetSceneMapInfoScRsp:                      "GetSceneMapInfoScRsp",
	SetClientPausedScRsp:                      "SetClientPausedScRsp",
	RefreshTriggerByClientScNotify:            "RefreshTriggerByClientScNotify",
	EnterSectionScRsp:                         "EnterSectionScRsp",
	GroupStateChangeScRsp:                     "GroupStateChangeScRsp",
	GetEnteredSceneScRsp:                      "GetEnteredSceneScRsp",
	SceneEntityMoveScNotify:                   "SceneEntityMoveScNotify",
	UpdateGroupPropertyCsReq:                  "UpdateGroupPropertyCsReq",
	RefreshTriggerByClientScRsp:               "RefreshTriggerByClientScRsp",
	DeactivateFarmElementScRsp:                "DeactivateFarmElementScRsp",
	LastSpringRefreshTimeNotify:               "LastSpringRefreshTimeNotify",
	StartCocoonStageScRsp:                     "StartCocoonStageScRsp",
	GameplayCounterRecoverScRsp:               "GameplayCounterRecoverScRsp",
	RefreshTriggerByClientCsReq:               "RefreshTriggerByClientCsReq",
	SyncEntityBuffChangeListScNotify:          "SyncEntityBuffChangeListScNotify",
	SavePointsInfoNotify:                      "SavePointsInfoNotify",
	GetUnlockTeleportScRsp:                    "GetUnlockTeleportScRsp",
	SceneEntityTeleportCsReq:                  "SceneEntityTeleportCsReq",
	SceneCastSkillMpUpdateScNotify:            "SceneCastSkillMpUpdateScNotify",
	SetTrainWorldIdScRsp:                      "SetTrainWorldIdScRsp",
	SceneEntityTeleportScRsp:                  "SceneEntityTeleportScRsp",
	SetTrainWorldIdCsReq:                      "SetTrainWorldIdCsReq",
	ActivateFarmElementScRsp:                  "ActivateFarmElementScRsp",
	ScenePlaneEventScNotify:                   "ScenePlaneEventScNotify",
	GetEnteredSceneCsReq:                      "GetEnteredSceneCsReq",
	EntityBindPropScRsp:                       "EntityBindPropScRsp",
	GetCurSceneInfoCsReq:                      "GetCurSceneInfoCsReq",
	GameplayCounterCountDownScRsp:             "GameplayCounterCountDownScRsp",
	ChangePropTimelineInfoScRsp:               "ChangePropTimelineInfoScRsp",
	SceneReviveAfterRebattleCsReq:             "SceneReviveAfterRebattleCsReq",
	GroupStateChangeScNotify:                  "GroupStateChangeScNotify",
	ChangePropTimelineInfoCsReq:               "ChangePropTimelineInfoCsReq",
	EnterSceneScRsp:                           "EnterSceneScRsp",
	SceneEnterStageCsReq:                      "SceneEnterStageCsReq",
	UpdateFloorSavedValueNotify:               "UpdateFloorSavedValueNotify",
	GetShopListScRsp:                          "GetShopListScRsp",
	TakeCityShopRewardScRsp:                   "TakeCityShopRewardScRsp",
	BuyGoodsScRsp:                             "BuyGoodsScRsp",
	BuyGoodsCsReq:                             "BuyGoodsCsReq",
	GetShopListCsReq:                          "GetShopListCsReq",
	TakeCityShopRewardCsReq:                   "TakeCityShopRewardCsReq",
	CityShopInfoScNotify:                      "CityShopInfoScNotify",
	GetTutorialScRsp:                          "GetTutorialScRsp",
	UnlockTutorialScRsp:                       "UnlockTutorialScRsp",
	GetTutorialGuideScRsp:                     "GetTutorialGuideScRsp",
	UnlockTutorialGuideScRsp:                  "UnlockTutorialGuideScRsp",
	GetTutorialGuideCsReq:                     "GetTutorialGuideCsReq",
	FinishTutorialGuideScRsp:                  "FinishTutorialGuideScRsp",
	GetTutorialCsReq:                          "GetTutorialCsReq",
	UnlockTutorialCsReq:                       "UnlockTutorialCsReq",
	FinishTutorialGuideCsReq:                  "FinishTutorialGuideCsReq",
	UnlockTutorialGuideCsReq:                  "UnlockTutorialGuideCsReq",
	FinishTutorialScRsp:                       "FinishTutorialScRsp",
	FinishTutorialCsReq:                       "FinishTutorialCsReq",
	GetChallengeScRsp:                         "GetChallengeScRsp",
	StartPartialChallengeCsReq:                "StartPartialChallengeCsReq",
	LeaveChallengeScRsp:                       "LeaveChallengeScRsp",
	ChallengeBossPhaseSettleNotify:            "ChallengeBossPhaseSettleNotify",
	StartChallengeScRsp:                       "StartChallengeScRsp",
	RestartChallengePhaseCsReq:                "RestartChallengePhaseCsReq",
	GetChallengeGroupStatisticsScRsp:          "GetChallengeGroupStatisticsScRsp",
	StartChallengeCsReq:                       "StartChallengeCsReq",
	ChallengeLineupNotify:                     "ChallengeLineupNotify",
	EnterChallengeNextPhaseCsReq:              "EnterChallengeNextPhaseCsReq",
	GetChallengeCsReq:                         "GetChallengeCsReq",
	TakeChallengeRewardCsReq:                  "TakeChallengeRewardCsReq",
	LeaveChallengeCsReq:                       "LeaveChallengeCsReq",
	GetCurChallengeScRsp:                      "GetCurChallengeScRsp",
	StartPartialChallengeScRsp:                "StartPartialChallengeScRsp",
	TakeChallengeRewardScRsp:                  "TakeChallengeRewardScRsp",
	ChallengeSettleNotify:                     "ChallengeSettleNotify",
	GetCurChallengeCsReq:                      "GetCurChallengeCsReq",
	EnterChallengeNextPhaseScRsp:              "EnterChallengeNextPhaseScRsp",
	GetChallengeGroupStatisticsCsReq:          "GetChallengeGroupStatisticsCsReq",
	RestartChallengePhaseScRsp:                "RestartChallengePhaseScRsp",
	NOOHANIHDOB:                               "NOOHANIHDOB",
	MILMMEFNLNJ:                               "MILMMEFNLNJ",
	BOBPIMEIKNI:                               "BOBPIMEIKNI",
	EJBIMMPCPIO:                               "EJBIMMPCPIO",
	NJKDJFMNKHG:                               "NJKDJFMNKHG",
	AAJKBEEBDAM:                               "AAJKBEEBDAM",
	HGILBKEAANN:                               "HGILBKEAANN",
	EBMIGHOFADN:                               "EBMIGHOFADN",
	IKIOKKGFNEG:                               "IKIOKKGFNEG",
	LMBFKNBBIAN:                               "LMBFKNBBIAN",
	DNCMGMKKEHP:                               "DNCMGMKKEHP",
	EDEFMJMJHDG:                               "EDEFMJMJHDG",
	AHIFJEGEGBH:                               "AHIFJEGEGBH",
	JKGANDLBMDP:                               "JKGANDLBMDP",
	IEJDCCKFGEC:                               "IEJDCCKFGEC",
	GHPMFEBLDAM:                               "GHPMFEBLDAM",
	AIMEOHGGGPH:                               "AIMEOHGGGPH",
	FLMKDFPNLNN:                               "FLMKDFPNLNN",
	DJHJNJIKKPI:                               "DJHJNJIKKPI",
	KAMDNPBKKEL:                               "KAMDNPBKKEL",
	IHCJAJGPJKK:                               "IHCJAJGPJKK",
	NGGMFNPNKEN:                               "NGGMFNPNKEN",
	HGJJMEKHJAB:                               "HGJJMEKHJAB",
	LOEBIPHOOHN:                               "LOEBIPHOOHN",
	NIJGLMIEFIG:                               "NIJGLMIEFIG",
	KGAECMDGBNH:                               "KGAECMDGBNH",
	BKHJBHHFIBE:                               "BKHJBHHFIBE",
	DDEOCJJMGEA:                               "DDEOCJJMGEA",
	CELPBBFEKOP:                               "CELPBBFEKOP",
	PFIFBAKBDJF:                               "PFIFBAKBDJF",
	LAFLMCHDLME:                               "LAFLMCHDLME",
	LKHCFFGAEKF:                               "LKHCFFGAEKF",
	BFKHHONBNMB:                               "BFKHHONBNMB",
	PIIIJLCPAML:                               "PIIIJLCPAML",
	LOKBANIDPPG:                               "LOKBANIDPPG",
	KDHBFOHEBPD:                               "KDHBFOHEBPD",
	IPFIDHAANHM:                               "IPFIDHAANHM",
	EKCDCNHJIBC:                               "EKCDCNHJIBC",
	NNAFACGKDOB:                               "NNAFACGKDOB",
	MBCOMIBJKID:                               "MBCOMIBJKID",
	GMAHMMHHMAB:                               "GMAHMMHHMAB",
	IJFGBBJNDMP:                               "IJFGBBJNDMP",
	MJMMOIAINAK:                               "MJMMOIAINAK",
	PMJFLCCJGNB:                               "PMJFLCCJGNB",
	JDIGBEDAFAK:                               "JDIGBEDAFAK",
	CFCMNGNHGFB:                               "CFCMNGNHGFB",
	FDDECGPJDKD:                               "FDDECGPJDKD",
	JGMGPJBILHJ:                               "JGMGPJBILHJ",
	CBEDDKNMNDO:                               "CBEDDKNMNDO",
	DFEOPEHIBPM:                               "DFEOPEHIBPM",
	BODKPNEFBBP:                               "BODKPNEFBBP",
	KBHGEGNHGPD:                               "KBHGEGNHGPD",
	DFNEHAIHGKK:                               "DFNEHAIHGKK",
	GetGachaInfoScRsp:                         "GetGachaInfoScRsp",
	GetGachaCeilingScRsp:                      "GetGachaCeilingScRsp",
	DoGachaScRsp:                              "DoGachaScRsp",
	ExchangeGachaCeilingScRsp:                 "ExchangeGachaCeilingScRsp",
	DoGachaCsReq:                              "DoGachaCsReq",
	GetGachaInfoCsReq:                         "GetGachaInfoCsReq",
	GetGachaCeilingCsReq:                      "GetGachaCeilingCsReq",
	ExchangeGachaCeilingCsReq:                 "ExchangeGachaCeilingCsReq",
	SetGachaDecideItemScRsp:                   "SetGachaDecideItemScRsp",
	SetGachaDecideItemCsReq:                   "SetGachaDecideItemCsReq",
	GetNpcTakenRewardScRsp:                    "GetNpcTakenRewardScRsp",
	GetFirstTalkNpcScRsp:                      "GetFirstTalkNpcScRsp",
	TakeTalkRewardScRsp:                       "TakeTalkRewardScRsp",
	FinishFirstTalkNpcScRsp:                   "FinishFirstTalkNpcScRsp",
	TakeTalkRewardCsReq:                       "TakeTalkRewardCsReq",
	GetFirstTalkByPerformanceNpcScRsp:         "GetFirstTalkByPerformanceNpcScRsp",
	GetNpcTakenRewardCsReq:                    "GetNpcTakenRewardCsReq",
	GetFirstTalkNpcCsReq:                      "GetFirstTalkNpcCsReq",
	FinishFirstTalkByPerformanceNpcScRsp:      "FinishFirstTalkByPerformanceNpcScRsp",
	GetFirstTalkByPerformanceNpcCsReq:         "GetFirstTalkByPerformanceNpcCsReq",
	FinishFirstTalkByPerformanceNpcCsReq:      "FinishFirstTalkByPerformanceNpcCsReq",
	FinishFirstTalkNpcCsReq:                   "FinishFirstTalkNpcCsReq",
	SelectInclinationTextScRsp:                "SelectInclinationTextScRsp",
	SelectInclinationTextCsReq:                "SelectInclinationTextCsReq",
	StartRaidScRsp:                            "StartRaidScRsp",
	DelSaveRaidScNotify:                       "DelSaveRaidScNotify",
	LeaveRaidScRsp:                            "LeaveRaidScRsp",
	TakeChallengeRaidRewardCsReq:              "TakeChallengeRaidRewardCsReq",
	GetAllSaveRaidScRsp:                       "GetAllSaveRaidScRsp",
	LeaveRaidCsReq:                            "LeaveRaidCsReq",
	GetRaidInfoScRsp:                          "GetRaidInfoScRsp",
	StartRaidCsReq:                            "StartRaidCsReq",
	GetSaveRaidCsReq:                          "GetSaveRaidCsReq",
	RaidInfoNotify:                            "RaidInfoNotify",
	SetClientRaidTargetCountScRsp:             "SetClientRaidTargetCountScRsp",
	GetRaidInfoCsReq:                          "GetRaidInfoCsReq",
	RaidKickByServerScNotify:                  "RaidKickByServerScNotify",
	SetClientRaidTargetCountCsReq:             "SetClientRaidTargetCountCsReq",
	GetSaveRaidScRsp:                          "GetSaveRaidScRsp",
	GetChallengeRaidInfoScRsp:                 "GetChallengeRaidInfoScRsp",
	ChallengeRaidNotify:                       "ChallengeRaidNotify",
	GetAllSaveRaidCsReq:                       "GetAllSaveRaidCsReq",
	TakeChallengeRaidRewardScRsp:              "TakeChallengeRaidRewardScRsp",
	GetArchiveDataScRsp:                       "GetArchiveDataScRsp",
	GetUpdatedArchiveDataScRsp:                "GetUpdatedArchiveDataScRsp",
	GetUpdatedArchiveDataCsReq:                "GetUpdatedArchiveDataCsReq",
	GetArchiveDataCsReq:                       "GetArchiveDataCsReq",
	GetBigDataAllRecommendScRsp:               "GetBigDataAllRecommendScRsp",
	GetBigDataRecommendScRsp:                  "GetBigDataRecommendScRsp",
	GetBigDataRecommendCsReq:                  "GetBigDataRecommendCsReq",
	GetBigDataAllRecommendCsReq:               "GetBigDataAllRecommendCsReq",
	GetExpeditionDataScRsp:                    "GetExpeditionDataScRsp",
	AcceptExpeditionScRsp:                     "AcceptExpeditionScRsp",
	AcceptMultipleExpeditionScRsp:             "AcceptMultipleExpeditionScRsp",
	CancelActivityExpeditionCsReq:             "CancelActivityExpeditionCsReq",
	TakeMultipleExpeditionRewardScRsp:         "TakeMultipleExpeditionRewardScRsp",
	AcceptActivityExpeditionScRsp:             "AcceptActivityExpeditionScRsp",
	CancelActivityExpeditionScRsp:             "CancelActivityExpeditionScRsp",
	AcceptActivityExpeditionCsReq:             "AcceptActivityExpeditionCsReq",
	GetExpeditionDataCsReq:                    "GetExpeditionDataCsReq",
	TakeExpeditionRewardScRsp:                 "TakeExpeditionRewardScRsp",
	TakeExpeditionRewardCsReq:                 "TakeExpeditionRewardCsReq",
	CancelExpeditionScRsp:                     "CancelExpeditionScRsp",
	TakeActivityExpeditionRewardScRsp:         "TakeActivityExpeditionRewardScRsp",
	CancelExpeditionCsReq:                     "CancelExpeditionCsReq",
	TakeMultipleActivityExpeditionRewardScRsp: "TakeMultipleActivityExpeditionRewardScRsp",
	TakeActivityExpeditionRewardCsReq:         "TakeActivityExpeditionRewardCsReq",
	ExpeditionDataChangeScNotify:              "ExpeditionDataChangeScNotify",
	AcceptExpeditionCsReq:                     "AcceptExpeditionCsReq",
	TakeMultipleActivityExpeditionRewardCsReq: "TakeMultipleActivityExpeditionRewardCsReq",
	EGKKAPICIMP:                               "EGKKAPICIMP",
	DAHLBDGMHCH:                               "DAHLBDGMHCH",
	DCAOKHOAHMD:                               "DCAOKHOAHMD",
	JMDPENNDDHA:                               "JMDPENNDDHA",
	AJDKJFFJAME:                               "AJDKJFFJAME",
	ADHBNIFDGLD:                               "ADHBNIFDGLD",
	ICIDMOINJOO:                               "ICIDMOINJOO",
	LBLDGLJFLMC:                               "LBLDGLJFLMC",
	MDNGDCFHDGG:                               "MDNGDCFHDGG",
	FNLLDOPIPHD:                               "FNLLDOPIPHD",
	AIEAADPOFKA:                               "AIEAADPOFKA",
	JGNEMNFKCGJ:                               "JGNEMNFKCGJ",
	MAAANABLHCF:                               "MAAANABLHCF",
	LPHFJLKINIO:                               "LPHFJLKINIO",
	FNIKJFLFMHP:                               "FNIKJFLFMHP",
	OMNDPBIJKEO:                               "OMNDPBIJKEO",
	EKAMFGAIBIB:                               "EKAMFGAIBIB",
	JKLCPKEJGNO:                               "JKLCPKEJGNO",
	LOFBEJOFPNB:                               "LOFBEJOFPNB",
	FNLGJFEOIMD:                               "FNLGJFEOIMD",
	CPOANNNHKGP:                               "CPOANNNHKGP",
	KEBECGMKLNM:                               "KEBECGMKLNM",
	HHMDIBHFIFG:                               "HHMDIBHFIFG",
	BKBJGLIOAJC:                               "BKBJGLIOAJC",
	JELFFFKNFCH:                               "JELFFFKNFCH",
	GGLHJMHLPOI:                               "GGLHJMHLPOI",
	CurTrialActivityScNotify:                  "CurTrialActivityScNotify",
	GetLoginActivityScRsp:                     "GetLoginActivityScRsp",
	SubmitMaterialSubmitActivityMaterialCsReq: "SubmitMaterialSubmitActivityMaterialCsReq",
	GetActivityScheduleConfigScRsp:            "GetActivityScheduleConfigScRsp",
	TakeLoginActivityRewardScRsp:              "TakeLoginActivityRewardScRsp",
	GetAvatarDeliverRewardActivityDataCsReq:   "GetAvatarDeliverRewardActivityDataCsReq",
	StartTrialActivityScRsp:                   "StartTrialActivityScRsp",
	LeaveTrialActivityCsReq:                   "LeaveTrialActivityCsReq",
	GetTrialActivityDataCsReq:                 "GetTrialActivityDataCsReq",
	TrialActivityDataChangeScNotify:           "TrialActivityDataChangeScNotify",
	TakeMaterialSubmitActivityRewardScRsp:     "TakeMaterialSubmitActivityRewardScRsp",
	GetMaterialSubmitActivityDataScRsp:        "GetMaterialSubmitActivityDataScRsp",
	TakeLoginActivityRewardCsReq:              "TakeLoginActivityRewardCsReq",
	TakeTrialActivityRewardScRsp:              "TakeTrialActivityRewardScRsp",
	LeaveTrialActivityScRsp:                   "LeaveTrialActivityScRsp",
	GetLoginActivityCsReq:                     "GetLoginActivityCsReq",
	AvatarDeliverRewardTakeRewardScRsp:        "AvatarDeliverRewardTakeRewardScRsp",
	GetActivityScheduleConfigCsReq:            "GetActivityScheduleConfigCsReq",
	AvatarDeliverRewardChooseAvatarScRsp:      "AvatarDeliverRewardChooseAvatarScRsp",
	TakeMaterialSubmitActivityRewardCsReq:     "TakeMaterialSubmitActivityRewardCsReq",
	TakeTrialActivityRewardCsReq:              "TakeTrialActivityRewardCsReq",
	SubmitMaterialSubmitActivityMaterialScRsp: "SubmitMaterialSubmitActivityMaterialScRsp",
	EnterTrialActivityStageScRsp:              "EnterTrialActivityStageScRsp",
	StartTrialActivityCsReq:                   "StartTrialActivityCsReq",
	GetAvatarDeliverRewardActivityDataScRsp:   "GetAvatarDeliverRewardActivityDataScRsp",
	AvatarDeliverRewardChooseAvatarCsReq:      "AvatarDeliverRewardChooseAvatarCsReq",
	GetMaterialSubmitActivityDataCsReq:        "GetMaterialSubmitActivityDataCsReq",
	AvatarDeliverRewardTakeRewardCsReq:        "AvatarDeliverRewardTakeRewardCsReq",
	GetTrialActivityDataScRsp:                 "GetTrialActivityDataScRsp",
	GetNpcMessageGroupScRsp:                   "GetNpcMessageGroupScRsp",
	FinishItemIdScRsp:                         "FinishItemIdScRsp",
	GetNpcStatusScRsp:                         "GetNpcStatusScRsp",
	FinishSectionIdScRsp:                      "FinishSectionIdScRsp",
	GetNpcStatusCsReq:                         "GetNpcStatusCsReq",
	GetMissionMessageInfoScRsp:                "GetMissionMessageInfoScRsp",
	GetNpcMessageGroupCsReq:                   "GetNpcMessageGroupCsReq",
	FinishItemIdCsReq:                         "FinishItemIdCsReq",
	GetMissionMessageInfoCsReq:                "GetMissionMessageInfoCsReq",
	FinishSectionIdCsReq:                      "FinishSectionIdCsReq",
	FinishPerformSectionIdScRsp:               "FinishPerformSectionIdScRsp",
	FinishPerformSectionIdCsReq:               "FinishPerformSectionIdCsReq",
	GetPlayerBoardDataScRsp:                   "GetPlayerBoardDataScRsp",
	SetDisplayAvatarScRsp:                     "SetDisplayAvatarScRsp",
	SetHeadIconScRsp:                          "SetHeadIconScRsp",
	SetIsDisplayAvatarInfoScRsp:               "SetIsDisplayAvatarInfoScRsp",
	SetHeadIconCsReq:                          "SetHeadIconCsReq",
	SetAssistAvatarCsReq:                      "SetAssistAvatarCsReq",
	GetPlayerBoardDataCsReq:                   "GetPlayerBoardDataCsReq",
	SetPersonalCardScRsp:                      "SetPersonalCardScRsp",
	SetDisplayAvatarCsReq:                     "SetDisplayAvatarCsReq",
	SetPersonalCardCsReq:                      "SetPersonalCardCsReq",
	SetSignatureScRsp:                         "SetSignatureScRsp",
	SetAssistAvatarScRsp:                      "SetAssistAvatarScRsp",
	SetIsDisplayAvatarInfoCsReq:               "SetIsDisplayAvatarInfoCsReq",
	SetSignatureCsReq:                         "SetSignatureCsReq",
	GetCurAssistCsReq:                         "GetCurAssistCsReq",
	GetFriendRecommendLineupDetailCsReq:       "GetFriendRecommendLineupDetailCsReq",
	CurAssistChangedNotify:                    "CurAssistChangedNotify",
	TakeAssistRewardCsReq:                     "TakeAssistRewardCsReq",
	SetFriendMarkScRsp:                        "SetFriendMarkScRsp",
	GetFriendListInfoScRsp:                    "GetFriendListInfoScRsp",
	GetFriendRecommendListInfoCsReq:           "GetFriendRecommendListInfoCsReq",
	GetFriendApplyListInfoScRsp:               "GetFriendApplyListInfoScRsp",
	SetAssistScRsp:                            "SetAssistScRsp",
	DeleteBlacklistScRsp:                      "DeleteBlacklistScRsp",
	GetAssistListCsReq:                        "GetAssistListCsReq",
	GetPlayerDetailInfoScRsp:                  "GetPlayerDetailInfoScRsp",
	GetFriendRecommendLineupDetailScRsp:       "GetFriendRecommendLineupDetailScRsp",
	GetAssistHistoryScRsp:                     "GetAssistHistoryScRsp",
	ApplyFriendScRsp:                          "ApplyFriendScRsp",
	SetFriendMarkCsReq:                        "SetFriendMarkCsReq",
	SearchPlayerCsReq:                         "SearchPlayerCsReq",
	GetFriendDevelopmentInfoScRsp:             "GetFriendDevelopmentInfoScRsp",
	GetPlatformPlayerInfoScRsp:                "GetPlatformPlayerInfoScRsp",
	GetFriendRecommendLineupCsReq:             "GetFriendRecommendLineupCsReq",
	GetFriendAssistListScRsp:                  "GetFriendAssistListScRsp",
	GetFriendLoginInfoCsReq:                   "GetFriendLoginInfoCsReq",
	SetFriendRemarkNameScRsp:                  "SetFriendRemarkNameScRsp",
	SetForbidOtherApplyFriendCsReq:            "SetForbidOtherApplyFriendCsReq",
	SyncAddBlacklistScNotify:                  "SyncAddBlacklistScNotify",
	GetPlayerDetailInfoCsReq:                  "GetPlayerDetailInfoCsReq",
	GetFriendBattleRecordDetailScRsp:          "GetFriendBattleRecordDetailScRsp",
	GetFriendRecommendLineupScRsp:             "GetFriendRecommendLineupScRsp",
	SyncHandleFriendScNotify:                  "SyncHandleFriendScNotify",
	ReportPlayerScRsp:                         "ReportPlayerScRsp",
	GetFriendListInfoCsReq:                    "GetFriendListInfoCsReq",
	GetCurAssistScRsp:                         "GetCurAssistScRsp",
	SyncDeleteFriendScNotify:                  "SyncDeleteFriendScNotify",
	GetFriendApplyListInfoCsReq:               "GetFriendApplyListInfoCsReq",
	DeleteFriendScRsp:                         "DeleteFriendScRsp",
	GetPlatformPlayerInfoCsReq:                "GetPlatformPlayerInfoCsReq",
	SearchPlayerScRsp:                         "SearchPlayerScRsp",
	HandleFriendScRsp:                         "HandleFriendScRsp",
	SetFriendRemarkNameCsReq:                  "SetFriendRemarkNameCsReq",
	GetFriendBattleRecordDetailCsReq:          "GetFriendBattleRecordDetailCsReq",
	GetFriendLoginInfoScRsp:                   "GetFriendLoginInfoScRsp",
	GetFriendRecommendListInfoScRsp:           "GetFriendRecommendListInfoScRsp",
	NewAssistHistoryNotify:                    "NewAssistHistoryNotify",
	SetAssistCsReq:                            "SetAssistCsReq",
	GetAssistListScRsp:                        "GetAssistListScRsp",
	DeleteFriendCsReq:                         "DeleteFriendCsReq",
	GetAssistHistoryCsReq:                     "GetAssistHistoryCsReq",
	GetFriendDevelopmentInfoCsReq:             "GetFriendDevelopmentInfoCsReq",
	AddBlacklistCsReq:                         "AddBlacklistCsReq",
	SetForbidOtherApplyFriendScRsp:            "SetForbidOtherApplyFriendScRsp",
	TakeAssistRewardScRsp:                     "TakeAssistRewardScRsp",
	ApplyFriendCsReq:                          "ApplyFriendCsReq",
	HandleFriendCsReq:                         "HandleFriendCsReq",
	DeleteBlacklistCsReq:                      "DeleteBlacklistCsReq",
	AddBlacklistScRsp:                         "AddBlacklistScRsp",
	SyncApplyFriendScNotify:                   "SyncApplyFriendScNotify",
	GetFriendAssistListCsReq:                  "GetFriendAssistListCsReq",
	ReportPlayerCsReq:                         "ReportPlayerCsReq",
	BuyBpLevelCsReq:                           "BuyBpLevelCsReq",
	TakeBpRewardCsReq:                         "TakeBpRewardCsReq",
	BattlePassInfoNotify:                      "BattlePassInfoNotify",
	TakeBpRewardScRsp:                         "TakeBpRewardScRsp",
	BuyBpLevelScRsp:                           "BuyBpLevelScRsp",
	TakeAllRewardScRsp:                        "TakeAllRewardScRsp",
	GetJukeboxDataScRsp:                       "GetJukeboxDataScRsp",
	UnlockBackGroundMusicScRsp:                "UnlockBackGroundMusicScRsp",
	PlayBackGroundMusicScRsp:                  "PlayBackGroundMusicScRsp",
	TrialBackGroundMusicScRsp:                 "TrialBackGroundMusicScRsp",
	PlayBackGroundMusicCsReq:                  "PlayBackGroundMusicCsReq",
	GetJukeboxDataCsReq:                       "GetJukeboxDataCsReq",
	TrialBackGroundMusicCsReq:                 "TrialBackGroundMusicCsReq",
	HHOEFAFIDKF:                               "HHOEFAFIDKF",
	GKEJEIJHOBM:                               "GKEJEIJHOBM",
	OFDJCKNCNJJ:                               "OFDJCKNCNJJ",
	LKNDAIONPDK:                               "LKNDAIONPDK",
	GOBECKLNBKI:                               "GOBECKLNBKI",
	PBINMHDELNE:                               "PBINMHDELNE",
	DAKJPCHPDNE:                               "DAKJPCHPDNE",
	NFHONPNFGFC:                               "NFHONPNFGFC",
	NGCKNKILNKH:                               "NGCKNKILNKH",
	NFABOHBKKPM:                               "NFABOHBKKPM",
	NKNILCNPOGJ:                               "NKNILCNPOGJ",
	HDFNBAECDLI:                               "HDFNBAECDLI",
	FPLLNDGMCEL:                               "FPLLNDGMCEL",
	DMBFMIKOFJP:                               "DMBFMIKOFJP",
	DIJOAOKEOGN:                               "DIJOAOKEOGN",
	MHAHKDDGBMG:                               "MHAHKDDGBMG",
	BBALCNDAHAG:                               "BBALCNDAHAG",
	DPLFMIDINFM:                               "DPLFMIDINFM",
	DCPMEDCNFNI:                               "DCPMEDCNFNI",
	NMNBBAKGBCO:                               "NMNBBAKGBCO",
	MCALEHFBCNI:                               "MCALEHFBCNI",
	NCJNACMOGDG:                               "NCJNACMOGDG",
	GHBLGNMMCKJ:                               "GHBLGNMMCKJ",
	TakeApRewardScRsp:                         "TakeApRewardScRsp",
	TakeAllApRewardCsReq:                      "TakeAllApRewardCsReq",
	GetDailyActiveInfoScRsp:                   "GetDailyActiveInfoScRsp",
	GetDailyActiveInfoCsReq:                   "GetDailyActiveInfoCsReq",
	TakeApRewardCsReq:                         "TakeApRewardCsReq",
	DailyActiveInfoNotify:                     "DailyActiveInfoNotify",
	TakeAllApRewardScRsp:                      "TakeAllApRewardScRsp",
	GetRndOptionScRsp:                         "GetRndOptionScRsp",
	DailyFirstMeetPamScRsp:                    "DailyFirstMeetPamScRsp",
	DailyFirstMeetPamCsReq:                    "DailyFirstMeetPamCsReq",
	GetRndOptionCsReq:                         "GetRndOptionCsReq",
	CBNCNBMAHHG:                               "CBNCNBMAHHG",
	PICNBMNMICK:                               "PICNBMNMICK",
	IALBBBDEMEN:                               "IALBBBDEMEN",
	GetFightActivityDataScRsp:                 "GetFightActivityDataScRsp",
	TakeFightActivityRewardCsReq:              "TakeFightActivityRewardCsReq",
	EnterFightActivityStageCsReq:              "EnterFightActivityStageCsReq",
	FightActivityDataChangeScNotify:           "FightActivityDataChangeScNotify",
	GetFightActivityDataCsReq:                 "GetFightActivityDataCsReq",
	EnterFightActivityStageScRsp:              "EnterFightActivityStageScRsp",
	TakeFightActivityRewardScRsp:              "TakeFightActivityRewardScRsp",
	NMCKJKNBLKD:                               "NMCKJKNBLKD",
	CPKIJGOKEEI:                               "CPKIJGOKEEI",
	HODDGPADLPG:                               "HODDGPADLPG",
	CAFBJEBFAIC:                               "CAFBJEBFAIC",
	OOIGFJFMMHN:                               "OOIGFJFMMHN",
	JDLHPKNAEPH:                               "JDLHPKNAEPH",
	CJELMCNKCGN:                               "CJELMCNKCGN",
	CMLJLFNKMBD:                               "CMLJLFNKMBD",
	ELMLHIMCCFL:                               "ELMLHIMCCFL",
	EPCJCLFDMPK:                               "EPCJCLFDMPK",
	AOPFAGDEFAN:                               "AOPFAGDEFAN",
	PALCBGGAIME:                               "PALCBGGAIME",
	DCHBPIPBGOG:                               "DCHBPIPBGOG",
	APANIBFAOED:                               "APANIBFAOED",
	TextJoinQueryCsReq:                        "TextJoinQueryCsReq",
	EDKHIIKKODJ:                               "EDKHIIKKODJ",
	SendMsgScRsp:                              "SendMsgScRsp",
	GetPrivateChatHistoryScRsp:                "GetPrivateChatHistoryScRsp",
	PrivateMsgOfflineUsersScNotify:            "PrivateMsgOfflineUsersScNotify",
	GetChatFriendHistoryScRsp:                 "GetChatFriendHistoryScRsp",
	RevcMsgScNotify:                           "RevcMsgScNotify",
	MarkChatEmojiScRsp:                        "MarkChatEmojiScRsp",
	SendMsgCsReq:                              "SendMsgCsReq",
	GetLoginChatInfoCsReq:                     "GetLoginChatInfoCsReq",
	GetPrivateChatHistoryCsReq:                "GetPrivateChatHistoryCsReq",
	BatchMarkChatEmojiScRsp:                   "BatchMarkChatEmojiScRsp",
	MarkChatEmojiCsReq:                        "MarkChatEmojiCsReq",
	GetLoginChatInfoScRsp:                     "GetLoginChatInfoScRsp",
	GetChatFriendHistoryCsReq:                 "GetChatFriendHistoryCsReq",
	GetChatEmojiListScRsp:                     "GetChatEmojiListScRsp",
	GetChatEmojiListCsReq:                     "GetChatEmojiListCsReq",
	AcceptedPamMissionExpireScRsp:             "AcceptedPamMissionExpireScRsp",
	SyncAcceptedPamMissionNotify:              "SyncAcceptedPamMissionNotify",
	AcceptedPamMissionExpireCsReq:             "AcceptedPamMissionExpireCsReq",
	MazeKillDirectCsReq:                       "MazeKillDirectCsReq",
	SwitchMascotUpdateScNotify:                "SwitchMascotUpdateScNotify",
	GetSwitchMascotDataCsReq:                  "GetSwitchMascotDataCsReq",
	ShareScRsp:                                "ShareScRsp",
	TakePictureScRsp:                          "TakePictureScRsp",
	DifficultyAdjustmentUpdateDataScRsp:       "DifficultyAdjustmentUpdateDataScRsp",
	GetGunPlayDataScRsp:                       "GetGunPlayDataScRsp",
	DifficultyAdjustmentGetDataCsReq:          "DifficultyAdjustmentGetDataCsReq",
	GetShareDataScRsp:                         "GetShareDataScRsp",
	CancelDirectDeliveryNoticeCsReq:           "CancelDirectDeliveryNoticeCsReq",
	UpdateGunPlayDataCsReq:                    "UpdateGunPlayDataCsReq",
	GetUnreleasedBlockInfoScRsp:               "GetUnreleasedBlockInfoScRsp",
	GetMovieRacingDataScRsp:                   "GetMovieRacingDataScRsp",
	GetOrigamiPropInfoScRsp:                   "GetOrigamiPropInfoScRsp",
	GetShareDataCsReq:                         "GetShareDataCsReq",
	SecurityReportScRsp:                       "SecurityReportScRsp",
	UpdateMovieRacingDataScRsp:                "UpdateMovieRacingDataScRsp",
	ShareCsReq:                                "ShareCsReq",
	MazeKillDirectScRsp:                       "MazeKillDirectScRsp",
	SubmitOrigamiItemCsReq:                    "SubmitOrigamiItemCsReq",
	TakePictureCsReq:                          "TakePictureCsReq",
	TriggerVoiceScRsp:                         "TriggerVoiceScRsp",
	GetUnreleasedBlockInfoCsReq:               "GetUnreleasedBlockInfoCsReq",
	UpdateGunPlayDataScRsp:                    "UpdateGunPlayDataScRsp",
	GetMovieRacingDataCsReq:                   "GetMovieRacingDataCsReq",
	CancelDirectDeliveryNoticeScRsp:           "CancelDirectDeliveryNoticeScRsp",
	DifficultyAdjustmentUpdateDataCsReq:       "DifficultyAdjustmentUpdateDataCsReq",
	DifficultyAdjustmentGetDataScRsp:          "DifficultyAdjustmentGetDataScRsp",
	TriggerVoiceCsReq:                         "TriggerVoiceCsReq",
	DirectDeliveryScNotify:                    "DirectDeliveryScNotify",
	SubmitOrigamiItemScRsp:                    "SubmitOrigamiItemScRsp",
	GetSwitchMascotDataScRsp:                  "GetSwitchMascotDataScRsp",
	CancelCacheNotifyScRsp:                    "CancelCacheNotifyScRsp",
	GetGunPlayDataCsReq:                       "GetGunPlayDataCsReq",
	GetOrigamiPropInfoCsReq:                   "GetOrigamiPropInfoCsReq",
	CancelCacheNotifyCsReq:                    "CancelCacheNotifyCsReq",
	UpdateMovieRacingDataCsReq:                "UpdateMovieRacingDataCsReq",
	BPNMEBKDJEA:                               "BPNMEBKDJEA",
	KLNKDCOGKAB:                               "KLNKDCOGKAB",
	JCMIIGMBKED:                               "JCMIIGMBKED",
	EMCHJNAOJML:                               "EMCHJNAOJML",
	JFPLIKOOFEH:                               "JFPLIKOOFEH",
	OHFGLCDJBAJ:                               "OHFGLCDJBAJ",
	DPCFNEKEAGF:                               "DPCFNEKEAGF",
	LJOKCKAGIPN:                               "LJOKCKAGIPN",
	GBOGGNAKMFG:                               "GBOGGNAKMFG",
	ILAFJJCACNA:                               "ILAFJJCACNA",
	GLMFFPNGOFA:                               "GLMFFPNGOFA",
	IHCMMGAEBJL:                               "IHCMMGAEBJL",
	NDBJEMIIBLF:                               "NDBJEMIIBLF",
	MAJGIAAMOPL:                               "MAJGIAAMOPL",
	OPGOBPLGBOI:                               "OPGOBPLGBOI",
	BBPHBNHNNKJ:                               "BBPHBNHNNKJ",
	LOELHDFPJKD:                               "LOELHDFPJKD",
	DCDHMPBNLLA:                               "DCDHMPBNLLA",
	ABEEIIHOPKJ:                               "ABEEIIHOPKJ",
	EOBDAMHNOCD:                               "EOBDAMHNOCD",
	JJBABPCHHFN:                               "JJBABPCHHFN",
	ODIJHINADIN:                               "ODIJHINADIN",
	PBGPJKJIDJE:                               "PBGPJKJIDJE",
	LLGKDECGEHJ:                               "LLGKDECGEHJ",
	LJMDJDFODKM:                               "LJMDJDFODKM",
	LKGGBAGPDNO:                               "LKGGBAGPDNO",
	OBABKHGHKFB:                               "OBABKHGHKFB",
	NMEHFCJAACK:                               "NMEHFCJAACK",
	EJPNNCNLONC:                               "EJPNNCNLONC",
	LPMBCEGLLEN:                               "LPMBCEGLLEN",
	HDPCKNIPILG:                               "HDPCKNIPILG",
	DELNPHOJHII:                               "DELNPHOJHII",
	FKIEEDNBPCO:                               "FKIEEDNBPCO",
	BCFIILJHALJ:                               "BCFIILJHALJ",
	LCLEAOGMNML:                               "LCLEAOGMNML",
	BCCMMOJGDPM:                               "BCCMMOJGDPM",
	AHJKNPINKIC:                               "AHJKNPINKIC",
	NDGODKMMMFJ:                               "NDGODKMMMFJ",
	BMJHOMAFFKP:                               "BMJHOMAFFKP",
	AKMJJOPBNBL:                               "AKMJJOPBNBL",
	HOHHDIDEEMO:                               "HOHHDIDEEMO",
	EODCBBDGFDA:                               "EODCBBDGFDA",
	IMLKDFCLBIL:                               "IMLKDFCLBIL",
	JNBBFOIODIA:                               "JNBBFOIODIA",
	BFHHOIOBLMI:                               "BFHHOIOBLMI",
	TreasureDungeonFinishScNotify:             "TreasureDungeonFinishScNotify",
	InteractTreasureDungeonGridCsReq:          "InteractTreasureDungeonGridCsReq",
	UseTreasureDungeonItemScRsp:               "UseTreasureDungeonItemScRsp",
	FightTreasureDungeonMonsterScRsp:          "FightTreasureDungeonMonsterScRsp",
	GetTreasureDungeonActivityDataScRsp:       "GetTreasureDungeonActivityDataScRsp",
	QuitTreasureDungeonScRsp:                  "QuitTreasureDungeonScRsp",
	TreasureDungeonDataScNotify:               "TreasureDungeonDataScNotify",
	OpenTreasureDungeonGridCsReq:              "OpenTreasureDungeonGridCsReq",
	EnterTreasureDungeonScRsp:                 "EnterTreasureDungeonScRsp",
	GetTreasureDungeonActivityDataCsReq:       "GetTreasureDungeonActivityDataCsReq",
	UseTreasureDungeonItemCsReq:               "UseTreasureDungeonItemCsReq",
	InteractTreasureDungeonGridScRsp:          "InteractTreasureDungeonGridScRsp",
	EnterTreasureDungeonCsReq:                 "EnterTreasureDungeonCsReq",
	OpenTreasureDungeonGridScRsp:              "OpenTreasureDungeonGridScRsp",
	FightTreasureDungeonMonsterCsReq:          "FightTreasureDungeonMonsterCsReq",
	QuitTreasureDungeonCsReq:                  "QuitTreasureDungeonCsReq",
	NGDGKNLHDBB:                               "NGDGKNLHDBB",
	EOENAEACHIG:                               "EOENAEACHIG",
	IPALIMDICEL:                               "IPALIMDICEL",
	BJJLCBBDHOF:                               "BJJLCBBDHOF",
	PBOMOMDNPNC:                               "PBOMOMDNPNC",
	DFHMNEKALKP:                               "DFHMNEKALKP",
	IFNGOOHFDNF:                               "IFNGOOHFDNF",
	HMHOKLIPDEN:                               "HMHOKLIPDEN",
	ADLCNMFDMHH:                               "ADLCNMFDMHH",
	FNDIILLENPF:                               "FNDIILLENPF",
	BCIEDHOKLJN:                               "BCIEDHOKLJN",
	PEGNHMLBKJK:                               "PEGNHMLBKJK",
	JLLFEEOBFIB:                               "JLLFEEOBFIB",
	GAIEHAKNMMF:                               "GAIEHAKNMMF",
	CKGJHJGAEPP:                               "CKGJHJGAEPP",
	CNOOEFBFLCL:                               "CNOOEFBFLCL",
	MAMDABMCIMP:                               "MAMDABMCIMP",
	PDPHCLPCNBP:                               "PDPHCLPCNBP",
	HPIDGCJDPEE:                               "HPIDGCJDPEE",
	GKFKFEJIPBL:                               "GKFKFEJIPBL",
	JGLPIMNCIOO:                               "JGLPIMNCIOO",
	GJDJCDBBHOE:                               "GJDJCDBBHOE",
	MCODBGAMGAE:                               "MCODBGAMGAE",
	CEEICJCFGGC:                               "CEEICJCFGGC",
	CJGHAPLBFAB:                               "CJGHAPLBFAB",
	NIBOCPDLGFG:                               "NIBOCPDLGFG",
	CJINJHFHOPB:                               "CJINJHFHOPB",
	AEDKLLDFCPN:                               "AEDKLLDFCPN",
	DOFPCPJLEEC:                               "DOFPCPJLEEC",
	LCFIJCKNMDB:                               "LCFIJCKNMDB",
	HMLPKFEOHKH:                               "HMLPKFEOHKH",
	HEDNPLNCICC:                               "HEDNPLNCICC",
	CHBOOADLNEA:                               "CHBOOADLNEA",
	OKOJAAKFFDB:                               "OKOJAAKFFDB",
	EIKFIEJKMNC:                               "EIKFIEJKMNC",
	HOLMFMJBKBK:                               "HOLMFMJBKBK",
	HKNKCNMBEAO:                               "HKNKCNMBEAO",
	HJBDDADIPNL:                               "HJBDDADIPNL",
	MGMDENGANGH:                               "MGMDENGANGH",
	OEODHMAFEBF:                               "OEODHMAFEBF",
	FLEGGMNPHJD:                               "FLEGGMNPHJD",
	EIOPHCKBEEP:                               "EIOPHCKBEEP",
	NPEKKFMMKHL:                               "NPEKKFMMKHL",
	LIHOHDABGEC:                               "LIHOHDABGEC",
	FBMPMKJCPKD:                               "FBMPMKJCPKD",
	DFJDECBJKMB:                               "DFJDECBJKMB",
	KIGDOGNLBEJ:                               "KIGDOGNLBEJ",
	PGHEIEEKELK:                               "PGHEIEEKELK",
	HIFEIMICDML:                               "HIFEIMICDML",
	OICMOEJBNJD:                               "OICMOEJBNJD",
	OJJADENFNDJ:                               "OJJADENFNDJ",
	MJFPNIKJEFI:                               "MJFPNIKJEFI",
	ICEBBMAKHIE:                               "ICEBBMAKHIE",
	CHJJOFPENLA:                               "CHJJOFPENLA",
	GetAetherDivideInfoScRsp:                  "GetAetherDivideInfoScRsp",
	HLBCKCDICAA:                               "HLBCKCDICAA",
	FMHGDKEFBBJ:                               "FMHGDKEFBBJ",
	GetAetherDivideInfoCsReq:                  "GetAetherDivideInfoCsReq",
	KEJCGNHOAML:                               "KEJCGNHOAML",
	DBJKDLIOIEO:                               "DBJKDLIOIEO",
	MHJBACDKLFE:                               "MHJBACDKLFE",
	LJNBHDOAOOC:                               "LJNBHDOAOOC",
	MHCGODLNNPJ:                               "MHCGODLNNPJ",
	LIHKFBKOGNB:                               "LIHKFBKOGNB",
	CJCJEFBPEBC:                               "CJCJEFBPEBC",
	ODIGIPEHBGG:                               "ODIGIPEHBGG",
	CMAFKHKKDNM:                               "CMAFKHKKDNM",
	ANDACJGGBEB:                               "ANDACJGGBEB",
	NOFNMHPMCJN:                               "NOFNMHPMCJN",
	KNCCOBPCFEL:                               "KNCCOBPCFEL",
	AEINLINFCDL:                               "AEINLINFCDL",
	OPEKCNGFDJO:                               "OPEKCNGFDJO",
	MPPINLPANKB:                               "MPPINLPANKB",
	GEJIONMNIKJ:                               "GEJIONMNIKJ",
	EDAGAAECHLK:                               "EDAGAAECHLK",
	FOJNPAEENOP:                               "FOJNPAEENOP",
	LAFNPEBLNDE:                               "LAFNPEBLNDE",
	DMLLGODHJPA:                               "DMLLGODHJPA",
	HEFCPAFNHHK:                               "HEFCPAFNHHK",
	DNHEIONHEIE:                               "DNHEIONHEIE",
	DIOHKAKCKLC:                               "DIOHKAKCKLC",
	KGJPEBCEPDL:                               "KGJPEBCEPDL",
	GetBenefitActivityInfoScRsp:               "GetBenefitActivityInfoScRsp",
	TakeBenefitActivityRewardScRsp:            "TakeBenefitActivityRewardScRsp",
	GetBenefitActivityInfoCsReq:               "GetBenefitActivityInfoCsReq",
	JoinBenefitActivityScRsp:                  "JoinBenefitActivityScRsp",
	JoinBenefitActivityCsReq:                  "JoinBenefitActivityCsReq",
	TakeBenefitActivityRewardCsReq:            "TakeBenefitActivityRewardCsReq",
	AHGMKOGAHGD:                               "AHGMKOGAHGD",
	HEHLNIMBFPM:                               "HEHLNIMBFPM",
	AFGANAFOOIJ:                               "AFGANAFOOIJ",
	NKGDPMONOPD:                               "NKGDPMONOPD",
	CEDICAHDICO:                               "CEDICAHDICO",
	PKIEOOOHPKG:                               "PKIEOOOHPKG",
	GetPhoneDataScRsp:                         "GetPhoneDataScRsp",
	SelectPhoneThemeCsReq:                     "SelectPhoneThemeCsReq",
	SelectChatBubbleScRsp:                     "SelectChatBubbleScRsp",
	UnlockPhoneThemeScNotify:                  "UnlockPhoneThemeScNotify",
	SelectChatBubbleCsReq:                     "SelectChatBubbleCsReq",
	GetPhoneDataCsReq:                         "GetPhoneDataCsReq",
	UnlockChatBubbleScNotify:                  "UnlockChatBubbleScNotify",
	UnlockPhoneCaseScNotify:                   "UnlockPhoneCaseScNotify",
	SelectPhoneThemeScRsp:                     "SelectPhoneThemeScRsp",
	SelectPhoneCaseScRsp:                      "SelectPhoneCaseScRsp",
	SelectPhoneCaseCsReq:                      "SelectPhoneCaseCsReq",
	CCKKHNGHAJM:                               "CCKKHNGHAJM",
	KNCFBDFLDFE:                               "KNCFBDFLDFE",
	FFOOFOHOOBD:                               "FFOOFOHOOBD",
	JNAKJHBCEGK:                               "JNAKJHBCEGK",
	PFLKCKJCDNP:                               "PFLKCKJCDNP",
	OLJNKMLOAAI:                               "OLJNKMLOAAI",
	KIDIMAGODDO:                               "KIDIMAGODDO",
	LILAOOMJGBO:                               "LILAOOMJGBO",
	HLMLPPGFCOK:                               "HLMLPPGFCOK",
	OOAKNKPPFJN:                               "OOAKNKPPFJN",
	OKIECEJGBNH:                               "OKIECEJGBNH",
	CCAMIJODAEO:                               "CCAMIJODAEO",
	MJKHAJMOFHD:                               "MJKHAJMOFHD",
	HGPGNDFJFNC:                               "HGPGNDFJFNC",
	FNKBMEGFCAO:                               "FNKBMEGFCAO",
	NGOPJFJGHPF:                               "NGOPJFJGHPF",
	DIPABHGIOHC:                               "DIPABHGIOHC",
	JOGMAKFDDDI:                               "JOGMAKFDDDI",
	KJLLINOKEEG:                               "KJLLINOKEEG",
	ICHJABJDILN:                               "ICHJABJDILN",
	IKLEMEDKBFA:                               "IKLEMEDKBFA",
	GAAKNMCLFDB:                               "GAAKNMCLFDB",
	CEEEEKKFEEN:                               "CEEEEKKFEEN",
	HCIMDMNDAJH:                               "HCIMDMNDAJH",
	PDLHODMBDGG:                               "PDLHODMBDGG",
	FDPNNADDHFB:                               "FDPNNADDHFB",
	BJCGPIDDKCO:                               "BJCGPIDDKCO",
	HJPPOLJOBDC:                               "HJPPOLJOBDC",
	DCEINFONEKB:                               "DCEINFONEKB",
	AJIDLGFBLBC:                               "AJIDLGFBLBC",
	LDJCMNILFIG:                               "LDJCMNILFIG",
	HJMLHAEOOGL:                               "HJMLHAEOOGL",
	HDELCMIKKHE:                               "HDELCMIKKHE",
	HFGIKBFCNHJ:                               "HFGIKBFCNHJ",
	OHKOBJKMJIF:                               "OHKOBJKMJIF",
	AKEADAIMKKB:                               "AKEADAIMKKB",
	KODBMIFICKL:                               "KODBMIFICKL",
	BGEPEAIEKBL:                               "BGEPEAIEKBL",
	MDFHEJDFPPF:                               "MDFHEJDFPPF",
	EMGANNEHNLJ:                               "EMGANNEHNLJ",
	PJBEEHEFGII:                               "PJBEEHEFGII",
	HJFPNHKKGEM:                               "HJFPNHKKGEM",
	BIGOGDODAFH:                               "BIGOGDODAFH",
	MIGIKKGEHAB:                               "MIGIKKGEHAB",
	AEJFOFIFGJJ:                               "AEJFOFIFGJJ",
	EKKDJDGIOCN:                               "EKKDJDGIOCN",
	FGBHBCKPEEC:                               "FGBHBCKPEEC",
	FMDBHHLPJLF:                               "FMDBHHLPJLF",
	DKMLFPLIOLM:                               "DKMLFPLIOLM",
	NCHANEEDHEB:                               "NCHANEEDHEB",
	OCAGICPAPNN:                               "OCAGICPAPNN",
	LEGNFIJHFJA:                               "LEGNFIJHFJA",
	IJIFCKAIFAM:                               "IJIFCKAIFAM",
	EGOEBBEKOAK:                               "EGOEBBEKOAK",
	AFBHPGLBLCH:                               "AFBHPGLBLCH",
	CEJFNFDKADC:                               "CEJFNFDKADC",
	PBJPGCCKDAA:                               "PBJPGCCKDAA",
	IFNPPMEGAGK:                               "IFNPPMEGAGK",
	LOKEPEOLDLK:                               "LOKEPEOLDLK",
	GFHPIELAEMD:                               "GFHPIELAEMD",
	BOPIBEAABHM:                               "BOPIBEAABHM",
	NIPHNIOEOOH:                               "NIPHNIOEOOH",
	MJGELCNMLIM:                               "MJGELCNMLIM",
	KKEJMJMKLBH:                               "KKEJMJMKLBH",
	LIHDJLDPBGO:                               "LIHDJLDPBGO",
	PHPGHELOJON:                               "PHPGHELOJON",
	MGCOCDHLOBL:                               "MGCOCDHLOBL",
	HJBINCHFMIJ:                               "HJBINCHFMIJ",
	ODOGHFDOAME:                               "ODOGHFDOAME",
	EBCGENEHPLO:                               "EBCGENEHPLO",
	NEEPIDOKIOE:                               "NEEPIDOKIOE",
	FOLENGGPDHF:                               "FOLENGGPDHF",
	BMEMEAIPICO:                               "BMEMEAIPICO",
	JOLMFHDNMJC:                               "JOLMFHDNMJC",
	AGMNCHOFGEP:                               "AGMNCHOFGEP",
	DHKBDLAKCJC:                               "DHKBDLAKCJC",
	OHPFCGMAOEK:                               "OHPFCGMAOEK",
	NJNDMAKLHGH:                               "NJNDMAKLHGH",
	LAMOKNACCFE:                               "LAMOKNACCFE",
	AIBMGOGLFLJ:                               "AIBMGOGLFLJ",
	CMJMFIKAJBG:                               "CMJMFIKAJBG",
	NHMDADNCDBN:                               "NHMDADNCDBN",
	NDMBKKGADNO:                               "NDMBKKGADNO",
	JMMJGOOJBEH:                               "JMMJGOOJBEH",
	CFIEHABLDOC:                               "CFIEHABLDOC",
	AEHCIIHAMBJ:                               "AEHCIIHAMBJ",
	BADNPCNJGOM:                               "BADNPCNJGOM",
	ODFBPOALLPD:                               "ODFBPOALLPD",
	LNDHKMMPBFJ:                               "LNDHKMMPBFJ",
	HCEEKIOAEKN:                               "HCEEKIOAEKN",
	JKOFLFFJDBM:                               "JKOFLFFJDBM",
	NELFLKMAAAG:                               "NELFLKMAAAG",
	FJGGNHODADK:                               "FJGGNHODADK",
	FFEBCHCIACF:                               "FFEBCHCIACF",
	MPJCFEFLGJK:                               "MPJCFEFLGJK",
	HPEJDAODMLN:                               "HPEJDAODMLN",
	JBENCNEAKPE:                               "JBENCNEAKPE",
	EMPIBLJGIOC:                               "EMPIBLJGIOC",
	LILKJLIGCEL:                               "LILKJLIGCEL",
	MKKJIHPPCNF:                               "MKKJIHPPCNF",
	CCBKNLDOLJH:                               "CCBKNLDOLJH",
	KCOMNODHCOD:                               "KCOMNODHCOD",
	CADKIDNGMCB:                               "CADKIDNGMCB",
	HNEFKBODKKN:                               "HNEFKBODKKN",
	IJNIPLKKIBO:                               "IJNIPLKKIBO",
	LBOFIPJNCAP:                               "LBOFIPJNCAP",
	DELPONBAIFF:                               "DELPONBAIFF",
	EAMNGABAKGB:                               "EAMNGABAKGB",
	BCHGLIJBDHA:                               "BCHGLIJBDHA",
	MOCFDOHFPCF:                               "MOCFDOHFPCF",
	EMDBEGPFMAO:                               "EMDBEGPFMAO",
	PMONFIOMJOC:                               "PMONFIOMJOC",
	MLPGILOAGLH:                               "MLPGILOAGLH",
	NJDDMAPMNKL:                               "NJDDMAPMNKL",
	NGPONCGIHAG:                               "NGPONCGIHAG",
	OEHGLLPFDFA:                               "OEHGLLPFDFA",
	JCIGNGGPCBN:                               "JCIGNGGPCBN",
	HIBPKPMLBBP:                               "HIBPKPMLBBP",
	OJBDHNMJPLH:                               "OJBDHNMJPLH",
	KENLCMJBHMC:                               "KENLCMJBHMC",
	MPLLHJCDKFN:                               "MPLLHJCDKFN",
	CBAFJJELCOI:                               "CBAFJJELCOI",
	IBDFBFGGGHM:                               "IBDFBFGGGHM",
	GLFANALPNBG:                               "GLFANALPNBG",
	FHKEEDOLOBO:                               "FHKEEDOLOBO",
	BJDHAPHDEHH:                               "BJDHAPHDEHH",
	DAMECIJDLOE:                               "DAMECIJDLOE",
	CFNGEEOFONA:                               "CFNGEEOFONA",
	ICPJOFDGLPC:                               "ICPJOFDGLPC",
	AFKBBNDLPBE:                               "AFKBBNDLPBE",
	AEBNKGPCHPC:                               "AEBNKGPCHPC",
	MBNABDOJJJN:                               "MBNABDOJJJN",
	EMILPJEHKKD:                               "EMILPJEHKKD",
	NAHGHNMDPKB:                               "NAHGHNMDPKB",
	AOMJBJFEEAA:                               "AOMJBJFEEAA",
	JIOFFGNHLLM:                               "JIOFFGNHLLM",
	AGLBFJIIDKJ:                               "AGLBFJIIDKJ",
	DIMMGDFHAED:                               "DIMMGDFHAED",
	EKMBHACDNCN:                               "EKMBHACDNCN",
	ANLNAJLMEOE:                               "ANLNAJLMEOE",
	BJAJDPCGDAE:                               "BJAJDPCGDAE",
	GGPBCFDAHGK:                               "GGPBCFDAHGK",
	GetBattleCollegeDataScRsp:                 "GetBattleCollegeDataScRsp",
	StartBattleCollegeCsReq:                   "StartBattleCollegeCsReq",
	BattleCollegeDataChangeScNotify:           "BattleCollegeDataChangeScNotify",
	GetBattleCollegeDataCsReq:                 "GetBattleCollegeDataCsReq",
	StartBattleCollegeScRsp:                   "StartBattleCollegeScRsp",
	HeliobusActivityDataScRsp:                 "HeliobusActivityDataScRsp",
	NGLFFFNCJLN:                               "NGLFFFNCJLN",
	LLMJHIHLBAD:                               "LLMJHIHLBAD",
	LFDIHDCPJCP:                               "LFDIHDCPJCP",
	ODELNDKFBKD:                               "ODELNDKFBKD",
	FHDCAFMEFLJ:                               "FHDCAFMEFLJ",
	MNILINDCMGJ:                               "MNILINDCMGJ",
	IJCKOIFEHGO:                               "IJCKOIFEHGO",
	HeliobusActivityDataCsReq:                 "HeliobusActivityDataCsReq",
	LGBFJCIBOIK:                               "LGBFJCIBOIK",
	KMHCOMICEIM:                               "KMHCOMICEIM",
	LCFNJLBHFBA:                               "LCFNJLBHFBA",
	CNFNJCECHGO:                               "CNFNJCECHGO",
	LCMIMFDLDIH:                               "LCMIMFDLDIH",
	HEONENLGLFB:                               "HEONENLGLFB",
	IDAEPKPFBEA:                               "IDAEPKPFBEA",
	DGMALFEMGCJ:                               "DGMALFEMGCJ",
	EIFABALINBA:                               "EIFABALINBA",
	OMMHMMGBOME:                               "OMMHMMGBOME",
	LHENFPKMHMK:                               "LHENFPKMHMK",
	LJFMPHJCKBG:                               "LJFMPHJCKBG",
	CDLFDJHKNLA:                               "CDLFDJHKNLA",
	DEMGAOOOFGJ:                               "DEMGAOOOFGJ",
	IDOGDCHCOMJ:                               "IDOGDCHCOMJ",
	JDFGHDDLNJE:                               "JDFGHDDLNJE",
	NPGOMPIDMFJ:                               "NPGOMPIDMFJ",
	DMFFKCBGKKJ:                               "DMFFKCBGKKJ",
	CNACENHJLOL:                               "CNACENHJLOL",
	NANEFGENLAO:                               "NANEFGENLAO",
	HPAFBMDJODD:                               "HPAFBMDJODD",
	MICGLCNILOL:                               "MICGLCNILOL",
	EGMPIICELAH:                               "EGMPIICELAH",
	MDLDAPDJNMK:                               "MDLDAPDJNMK",
	DPOLNPLHIAA:                               "DPOLNPLHIAA",
	JLKBPPLJKIJ:                               "JLKBPPLJKIJ",
	MMPJPKEODEL:                               "MMPJPKEODEL",
	EBCACFHPPCF:                               "EBCACFHPPCF",
	EMOPJEGPHGK:                               "EMOPJEGPHGK",
	GHCCOIMOGJB:                               "GHCCOIMOGJB",
	NOEBCPDGJAC:                               "NOEBCPDGJAC",
	PANODPAJMOJ:                               "PANODPAJMOJ",
	PJCCJAFEPMP:                               "PJCCJAFEPMP",
	OIFFDIOMIIF:                               "OIFFDIOMIIF",
	JDBGCNBDLLO:                               "JDBGCNBDLLO",
	PAMKNHOKHDL:                               "PAMKNHOKHDL",
	ADMNIBIJHFL:                               "ADMNIBIJHFL",
	HJFJAAFODAC:                               "HJFJAAFODAC",
	FDCNFHLJEJH:                               "FDCNFHLJEJH",
	FJFKPKDAKGF:                               "FJFKPKDAKGF",
	FMLLCFAEGKF:                               "FMLLCFAEGKF",
	DHKKJDJOKMH:                               "DHKKJDJOKMH",
	PKKGGNJICEO:                               "PKKGGNJICEO",
	CJMDEHEGBDB:                               "CJMDEHEGBDB",
	NDBCPOEKMAP:                               "NDBCPOEKMAP",
	CCFKEPHJECI:                               "CCFKEPHJECI",
	AACAGDMNDIA:                               "AACAGDMNDIA",
	CKNFINDGJON:                               "CKNFINDGJON",
	PGPKFJMBNBG:                               "PGPKFJMBNBG",
	JNAFHOCCEOA:                               "JNAFHOCCEOA",
	MKIBJDIGCOJ:                               "MKIBJDIGCOJ",
	DDJJOEOEAFE:                               "DDJJOEOEAFE",
	JAACIPEBEFB:                               "JAACIPEBEFB",
	GGBHOHNPHBO:                               "GGBHOHNPHBO",
	PCFFAIPODFI:                               "PCFFAIPODFI",
	JBAGBPPIHFC:                               "JBAGBPPIHFC",
	IHAKOOAGOHF:                               "IHAKOOAGOHF",
	JBJBLDDKCFF:                               "JBJBLDDKCFF",
	HLBDHKLBFKL:                               "HLBDHKLBFKL",
	IKAIOFPBHJM:                               "IKAIOFPBHJM",
	LONGMLLNJHJ:                               "LONGMLLNJHJ",
	AOOCNHIDGNN:                               "AOOCNHIDGNN",
	CEMEPHAFBEN:                               "CEMEPHAFBEN",
	PCHHLGNHFJL:                               "PCHHLGNHFJL",
	OGBGKDNMJJN:                               "OGBGKDNMJJN",
	GOPCELBLMEK:                               "GOPCELBLMEK",
	OMMPOGNONLL:                               "OMMPOGNONLL",
	NLEOJELEADF:                               "NLEOJELEADF",
	OHCIMEBNOOA:                               "OHCIMEBNOOA",
	BNLFGHMPGLL:                               "BNLFGHMPGLL",
	BEBCPIBJNOI:                               "BEBCPIBJNOI",
	HPBOKBBJOHD:                               "HPBOKBBJOHD",
	KFAPGHFPOAM:                               "KFAPGHFPOAM",
	KPFIALIIHLK:                               "KPFIALIIHLK",
	GIEMCFEHCII:                               "GIEMCFEHCII",
	KCEJGCBBPPE:                               "KCEJGCBBPPE",
	OMLIBKJMKKC:                               "OMLIBKJMKKC",
	FPBFPJJHEAA:                               "FPBFPJJHEAA",
	OAKPNPEKJFF:                               "OAKPNPEKJFF",
	HIECGAOJIHM:                               "HIECGAOJIHM",
	DNJFHOHBEIH:                               "DNJFHOHBEIH",
	CMMAFJGOEBO:                               "CMMAFJGOEBO",
	DGJOJAEFENJ:                               "DGJOJAEFENJ",
	LLOHGKBAPOE:                               "LLOHGKBAPOE",
	KBLGBOKEIJE:                               "KBLGBOKEIJE",
	RogueTournGetCurRogueCocoonInfoCsReq:      "RogueTournGetCurRogueCocoonInfoCsReq",
	FCHBLIEKDIO:                               "FCHBLIEKDIO",
	MMBHMGMADNP:                               "MMBHMGMADNP",
	AKBJCJBIHLC:                               "AKBJCJBIHLC",
	PIBAKIJCEOL:                               "PIBAKIJCEOL",
	CFHDDAAJFIO:                               "CFHDDAAJFIO",
	OJLEHAPCALO:                               "OJLEHAPCALO",
	PAOBPIIDBDB:                               "PAOBPIIDBDB",
	KLPKPGEPOPL:                               "KLPKPGEPOPL",
	EFKCPPCLINI:                               "EFKCPPCLINI",
	OOACPALGDLJ:                               "OOACPALGDLJ",
	DLIFNCLEKHL:                               "DLIFNCLEKHL",
	MDNBEPPCHCJ:                               "MDNBEPPCHCJ",
	GBKEEFDDELM:                               "GBKEEFDDELM",
	OOILNIEOEHE:                               "OOILNIEOEHE",
	BNGDJIDDHNF:                               "BNGDJIDDHNF",
	LNLDMFOFFNM:                               "LNLDMFOFFNM",
	LKCMNAAIMMM:                               "LKCMNAAIMMM",
	BOKJMAGOFGA:                               "BOKJMAGOFGA",
	JIFGEAGMFGG:                               "JIFGEAGMFGG",
	MLMDECIBOEB:                               "MLMDECIBOEB",
	HJPIIBCMGGC:                               "HJPIIBCMGGC",
	KPBBMANACIC:                               "KPBBMANACIC",
	RogueTournGetCurRogueCocoonInfoScRsp:      "RogueTournGetCurRogueCocoonInfoScRsp",
	MBAJJFBEGMI:                               "MBAJJFBEGMI",
	KAADNHDENPI:                               "KAADNHDENPI",
	KAPJNAKMKOB:                               "KAPJNAKMKOB",
	BFDADNAEHBG:                               "BFDADNAEHBG",
	EABGBGMMPKN:                               "EABGBGMMPKN",
	PDBEEBMJJDB:                               "PDBEEBMJJDB",
	DPFFEELFABM:                               "DPFFEELFABM",
	CCCJDJCBJLO:                               "CCCJDJCBJLO",
	HPLHFAIIDEO:                               "HPLHFAIIDEO",
	IEGIMEHHLKH:                               "IEGIMEHHLKH",
	DDPEFJCMLJL:                               "DDPEFJCMLJL",
	IDJKKNKJGFP:                               "IDJKKNKJGFP",
	BOAOLFBJAGJ:                               "BOAOLFBJAGJ",
	OFHDPLMJPCA:                               "OFHDPLMJPCA",
	GetAllServerPrefsDataScRsp:                "GetAllServerPrefsDataScRsp",
	UpdateServerPrefsDataScRsp:                "UpdateServerPrefsDataScRsp",
	GetServerPrefsDataScRsp:                   "GetServerPrefsDataScRsp",
	GetAllServerPrefsDataCsReq:                "GetAllServerPrefsDataCsReq",
	UpdateServerPrefsDataCsReq:                "UpdateServerPrefsDataCsReq",
	GetStoryLineInfoScRsp:                     "GetStoryLineInfoScRsp",
	ChangeStoryLineFinishScNotify:             "ChangeStoryLineFinishScNotify",
	StoryLineInfoScNotify:                     "StoryLineInfoScNotify",
	GetStoryLineInfoCsReq:                     "GetStoryLineInfoCsReq",
	StoryLineTrialAvatarChangeScNotify:        "StoryLineTrialAvatarChangeScNotify",
	GetHeartDialInfoScRsp:                     "GetHeartDialInfoScRsp",
	SubmitEmotionItemScRsp:                    "SubmitEmotionItemScRsp",
	ChangeScriptEmotionScRsp:                  "ChangeScriptEmotionScRsp",
	FinishEmotionDialoguePerformanceScRsp:     "FinishEmotionDialoguePerformanceScRsp",
	ChangeScriptEmotionCsReq:                  "ChangeScriptEmotionCsReq",
	GetHeartDialInfoCsReq:                     "GetHeartDialInfoCsReq",
	HeartDialTraceScriptScRsp:                 "HeartDialTraceScriptScRsp",
	FinishEmotionDialoguePerformanceCsReq:     "FinishEmotionDialoguePerformanceCsReq",
	HeartDialTraceScriptCsReq:                 "HeartDialTraceScriptCsReq",
	HeartDialScriptChangeScNotify:             "HeartDialScriptChangeScNotify",
	ODNHKDDNLPB:                               "ODNHKDDNLPB",
	KALFPKOLCMC:                               "KALFPKOLCMC",
	HELMGHHLLMB:                               "HELMGHHLLMB",
	MBCONMAHPGC:                               "MBCONMAHPGC",
	PDNGJBANLEH:                               "PDNGJBANLEH",
	IEOHIBMALDB:                               "IEOHIBMALDB",
	AKOOJPGGIID:                               "AKOOJPGGIID",
	PIIHEFOPBGL:                               "PIIHEFOPBGL",
	LOEJLDBMELC:                               "LOEJLDBMELC",
	GMOJKLOLCMI:                               "GMOJKLOLCMI",
	KNLCKNHNEIO:                               "KNLCKNHNEIO",
	LJAEKJBGAJH:                               "LJAEKJBGAJH",
	LJGHHMAFMJC:                               "LJGHHMAFMJC",
	PDFNJCMJFFE:                               "PDFNJCMJFFE",
	AADDJICOPGG:                               "AADDJICOPGG",
	HPNFJEGCFDM:                               "HPNFJEGCFDM",
	CPBMOIIPKMP:                               "CPBMOIIPKMP",
	GNGFHKPCGJO:                               "GNGFHKPCGJO",
	BFPILCMLFDN:                               "BFPILCMLFDN",
	KDAAEBDKEJG:                               "KDAAEBDKEJG",
	JHCOIGILILA:                               "JHCOIGILILA",
	GCAOGCEIAFP:                               "GCAOGCEIAFP",
	JJIBHDBEHJE:                               "JJIBHDBEHJE",
	ABOADJAAGBL:                               "ABOADJAAGBL",
	DPOGEMAPOKL:                               "DPOGEMAPOKL",
	BHHHPJOIJPH:                               "BHHHPJOIJPH",
	DHFHJJKPPPD:                               "DHFHJJKPPPD",
	ICIEJEBBKJE:                               "ICIEJEBBKJE",
	MCHKEHCFKKO:                               "MCHKEHCFKKO",
	CBIKDGJILJF:                               "CBIKDGJILJF",
	AOCPKLPBLEF:                               "AOCPKLPBLEF",
	CKGENIOHBPL:                               "CKGENIOHBPL",
	KDOMNBDDOBE:                               "KDOMNBDDOBE",
	HBDNGKNAKBL:                               "HBDNGKNAKBL",
	LBOLCAABCCG:                               "LBOLCAABCCG",
	JOABMHCDHDL:                               "JOABMHCDHDL",
	EnterEraFlipperRegionCsReq:                "EnterEraFlipperRegionCsReq",
	OJOBJIHLCAK:                               "OJOBJIHLCAK",
	BDAGNLNGHIC:                               "BDAGNLNGHIC",
	ChangeEraFlipperDataScRsp:                 "ChangeEraFlipperDataScRsp",
	PKBPGMGMDAC:                               "PKBPGMGMDAC",
	GetEraFlipperDataScRsp:                    "GetEraFlipperDataScRsp",
	AKHJLLNNAJA:                               "AKHJLLNNAJA",
	IHEOCPCAMBF:                               "IHEOCPCAMBF",
	BGDHCMBNNFC:                               "BGDHCMBNNFC",
	KGIPBNHOJLG:                               "KGIPBNHOJLG",
	CCBCMAOCEDG:                               "CCBCMAOCEDG",
	OCPJHDOFIMC:                               "OCPJHDOFIMC",
	PFFDLKILHFM:                               "PFFDLKILHFM",
	CHJANNPMLEO:                               "CHJANNPMLEO",
	LHBCEGEFIAC:                               "LHBCEGEFIAC",
	LFDJEIEKICL:                               "LFDJEIEKICL",
	MAFFBCDOIEM:                               "MAFFBCDOIEM",
	JEEHLBCGFJK:                               "JEEHLBCGFJK",
	CDIEDODAHDO:                               "CDIEDODAHDO",
	MFPFDOKCIPH:                               "MFPFDOKCIPH",
	KFNDBFAEMJN:                               "KFNDBFAEMJN",
	OHHPBOEJLLC:                               "OHHPBOEJLLC",
	DAGIOJOBGEE:                               "DAGIOJOBGEE",
	NFELMOCJCMM:                               "NFELMOCJCMM",
	MIBIMKEPCME:                               "MIBIMKEPCME",
	KJIEOKEFPAF:                               "KJIEOKEFPAF",
	LOPHALHNDHF:                               "LOPHALHNDHF",
	PIAMJEHGGHH:                               "PIAMJEHGGHH",
	EGOCFONOBKF:                               "EGOCFONOBKF",
	LJAEKBDJJFM:                               "LJAEKBDJJFM",
	DNBHCGKJKAH:                               "DNBHCGKJKAH",
	OGHNACAMIEC:                               "OGHNACAMIEC",
	GBBMKCNJBOM:                               "GBBMKCNJBOM",
	PFLIEDNIGPA:                               "PFLIEDNIGPA",
	EnterMapRotationRegionScRsp:               "EnterMapRotationRegionScRsp",
	RemoveRotaterScRsp:                        "RemoveRotaterScRsp",
	DeployRotaterScRsp:                        "DeployRotaterScRsp",
	InteractChargerScRsp:                      "InteractChargerScRsp",
	RotateMapScRsp:                            "RotateMapScRsp",
	RemoveRotaterCsReq:                        "RemoveRotaterCsReq",
	InteractChargerCsReq:                      "InteractChargerCsReq",
	GetMapRotationDataScRsp:                   "GetMapRotationDataScRsp",
	EnterMapRotationRegionCsReq:               "EnterMapRotationRegionCsReq",
	LeaveMapRotationRegionScNotify:            "LeaveMapRotationRegionScNotify",
	DeployRotaterCsReq:                        "DeployRotaterCsReq",
	ResetMapRotationRegionScRsp:               "ResetMapRotationRegionScRsp",
	GetMapRotationDataCsReq:                   "GetMapRotationDataCsReq",
	UpdateRotaterScNotify:                     "UpdateRotaterScNotify",
	ResetMapRotationRegionCsReq:               "ResetMapRotationRegionCsReq",
	UpdateEnergyScNotify:                      "UpdateEnergyScNotify",
	RotateMapCsReq:                            "RotateMapCsReq",
	LeaveMapRotationRegionScRsp:               "LeaveMapRotationRegionScRsp",
	UpdateMapRotationDataScNotify:             "UpdateMapRotationDataScNotify",
	LeaveMapRotationRegionCsReq:               "LeaveMapRotationRegionCsReq",
	TakeRollShopRewardScRsp:                   "TakeRollShopRewardScRsp",
	DoGachaInRollShopScRsp:                    "DoGachaInRollShopScRsp",
	GetRollShopInfoScRsp:                      "GetRollShopInfoScRsp",
	DoGachaInRollShopCsReq:                    "DoGachaInRollShopCsReq",
	GetRollShopInfoCsReq:                      "GetRollShopInfoCsReq",
	TakeRollShopRewardCsReq:                   "TakeRollShopRewardCsReq",
	OfferingInfoScNotify:                      "OfferingInfoScNotify",
	TakeOfferingRewardScRsp:                   "TakeOfferingRewardScRsp",
	SubmitOfferingItemScRsp:                   "SubmitOfferingItemScRsp",
	GetOfferingInfoScRsp:                      "GetOfferingInfoScRsp",
	SubmitOfferingItemCsReq:                   "SubmitOfferingItemCsReq",
	GetOfferingInfoCsReq:                      "GetOfferingInfoCsReq",
	TakeOfferingRewardCsReq:                   "TakeOfferingRewardCsReq",
	FBNKKFNMLOM:                               "FBNKKFNMLOM",
	RaidCollectionDataScRsp:                   "RaidCollectionDataScRsp",
	OOIBELOGDIA:                               "OOIBELOGDIA",
	RaidCollectionDataCsReq:                   "RaidCollectionDataCsReq",
	IIKBHNKPJFO:                               "IIKBHNKPJFO",
	BLHIJCEDNNE:                               "BLHIJCEDNNE",
	JEPAGHNAPCC:                               "JEPAGHNAPCC",
	GALFJGOGDHL:                               "GALFJGOGDHL",
	IEEJNPGCEGA:                               "IEEJNPGCEGA",
	LJDGJLPLHIH:                               "LJDGJLPLHIH",
	JPIKKKDADND:                               "JPIKKKDADND",
	JADFJJNOAAB:                               "JADFJJNOAAB",
	KADIILDGAPD:                               "KADIILDGAPD",
	NPOPGIHLDNO:                               "NPOPGIHLDNO",
	AKLPOCNEDDJ:                               "AKLPOCNEDDJ",
	MDDJJLJFACE:                               "MDDJJLJFACE",
	KMCHEILEKKI:                               "KMCHEILEKKI",
	OPHHNPGIKNK:                               "OPHHNPGIKNK",
	OGDMBJBOBDN:                               "OGDMBJBOBDN",
	LFHDJPDIHDO:                               "LFHDJPDIHDO",
	MFIEPPHJHBC:                               "MFIEPPHJHBC",
	FEPPFECHKPB:                               "FEPPFECHKPB",
	IMMCALOKIJN:                               "IMMCALOKIJN",
	NLBILPPLBFL:                               "NLBILPPLBFL",
	HAIBDDPANJJ:                               "HAIBDDPANJJ",
	GBNCOMHBFJA:                               "GBNCOMHBFJA",
	IAACILJIODI:                               "IAACILJIODI",
	FPKLBDNAPEG:                               "FPKLBDNAPEG",
	MGMACJHPELD:                               "MGMACJHPELD",
	JFGJMCCOPJF:                               "JFGJMCCOPJF",
	KGHHKLNJANA:                               "KGHHKLNJANA",
	CKCIPGPKAHB:                               "CKCIPGPKAHB",
	NJPCCBNPMEP:                               "NJPCCBNPMEP",
	DLEOPIFLMAO:                               "DLEOPIFLMAO",
	MEKBAOOIFOK:                               "MEKBAOOIFOK",
	GMICFOLDOPA:                               "GMICFOLDOPA",
	HJDFOHNJPPH:                               "HJDFOHNJPPH",
	FHLDKLFKEPG:                               "FHLDKLFKEPG",
	AGJIJJBBDKG:                               "AGJIJJBBDKG",
	JFONBMAAFEK:                               "JFONBMAAFEK",
	CILGFNEGEHD:                               "CILGFNEGEHD",
	LJIDIFFPAKL:                               "LJIDIFFPAKL",
	JFEEFPCMBPK:                               "JFEEFPCMBPK",
	JMFGGMFHEHJ:                               "JMFGGMFHEHJ",
	HILFAIHNGKD:                               "HILFAIHNGKD",
	DGEDOPCNLME:                               "DGEDOPCNLME",
	HNPCLJOKJII:                               "HNPCLJOKJII",
	ENEAMLFNHPE:                               "ENEAMLFNHPE",
	CCANODEDLHB:                               "CCANODEDLHB",
	OKGDEJIGADE:                               "OKGDEJIGADE",
	EJPBCEIBNJI:                               "EJPBCEIBNJI",
	JFJGHAFDFJE:                               "JFJGHAFDFJE",
	GCCMLJNDNFH:                               "GCCMLJNDNFH",
	FJOBKKKDMNK:                               "FJOBKKKDMNK",
	GDKGEAJPECK:                               "GDKGEAJPECK",
	LJPJGDEGOMH:                               "LJPJGDEGOMH",
	CMDMGBPBAID:                               "CMDMGBPBAID",
	EKIOMGCIINK:                               "EKIOMGCIINK",
	GDBCMFABPBJ:                               "GDBCMFABPBJ",
	NOHGHMMFBLC:                               "NOHGHMMFBLC",
	MKLFMDBBAON:                               "MKLFMDBBAON",
	MMJEHAILNGE:                               "MMJEHAILNGE",
	MAOFJEEAPJO:                               "MAOFJEEAPJO",
	IFEBFONCHCO:                               "IFEBFONCHCO",
	FGOIFCABMMG:                               "FGOIFCABMMG",
	GFJGFKCGLKM:                               "GFJGFKCGLKM",
	KGNHCCPMBLH:                               "KGNHCCPMBLH",
	KMBIPNPCADD:                               "KMBIPNPCADD",
	JAJOKJHMGBJ:                               "JAJOKJHMGBJ",
	OAFHJFCNGKJ:                               "OAFHJFCNGKJ",
	MEJLOBJPIJN:                               "MEJLOBJPIJN",
	IJMCPCMPJME:                               "IJMCPCMPJME",
	CFFJCAECFMD:                               "CFFJCAECFMD",
	BJOLJJALLLL:                               "BJOLJJALLLL",
	DJDOLLMENFG:                               "DJDOLLMENFG",
	KGGEIGMPLNL:                               "KGGEIGMPLNL",
	KAGABFDOIAO:                               "KAGABFDOIAO",
	JKGPHJNFDIA:                               "JKGPHJNFDIA",
	DPCPMOLHGME:                               "DPCPMOLHGME",
	KOIEKODNAGI:                               "KOIEKODNAGI",
	KOEBOMPEPJK:                               "KOEBOMPEPJK",
	PCFGODKODHM:                               "PCFGODKODHM",
	AECJKLDGHJK:                               "AECJKLDGHJK",
	ANHCFIPGLDE:                               "ANHCFIPGLDE",
	OHGHIBHDKLD:                               "OHGHIBHDKLD",
	JCAJLACNKDM:                               "JCAJLACNKDM",
	BGFOECMPAGB:                               "BGFOECMPAGB",
	LOEKMDFOIOF:                               "LOEKMDFOIOF",
	PCNKHELGLHJ:                               "PCNKHELGLHJ",
	BAOCKGBCLID:                               "BAOCKGBCLID",
	OFMKJHFPBNK:                               "OFMKJHFPBNK",
	KKDNFLDMGLM:                               "KKDNFLDMGLM",
	LFKLOAPMFKP:                               "LFKLOAPMFKP",
	BBPBBBDDKFI:                               "BBPBBBDDKFI",
	MNCBHMNLIFG:                               "MNCBHMNLIFG",
	DABEMHCDIEC:                               "DABEMHCDIEC",
	OBEIGOMKPFB:                               "OBEIGOMKPFB",
	MJPJNDCKOPF:                               "MJPJNDCKOPF",
	ANLCHKLHCIE:                               "ANLCHKLHCIE",
	LFMGJLCBAFB:                               "LFMGJLCBAFB",
	PJGMMJNDHMJ:                               "PJGMMJNDHMJ",
	GCPJPEJMDAP:                               "GCPJPEJMDAP",
	PKBAGKLFFGK:                               "PKBAGKLFFGK",
	EDIDBDGJIAF:                               "EDIDBDGJIAF",
	DCKFGKKDKEB:                               "DCKFGKKDKEB",
	FICMEMECAJK:                               "FICMEMECAJK",
	GPKACHPPGKI:                               "GPKACHPPGKI",
	AFPAEBGOLHG:                               "AFPAEBGOLHG",
	GLDPEJDGFEN:                               "GLDPEJDGFEN",
	HHFJJDMGIIP:                               "HHFJJDMGIIP",
	EPLPNAPFBHM:                               "EPLPNAPFBHM",
	DDBDPDMFLME:                               "DDBDPDMFLME",
	GetEraFlipperDataCsReq:                    "GetEraFlipperDataCsReq",
	ResetEraFlipperDataCsReq:                  "ResetEraFlipperDataCsReq",
	PEJJABDLOIH:                               "PEJJABDLOIH",
	EAOKBOCDHHG:                               "EAOKBOCDHHG",
	KJLCMECMHHE:                               "KJLCMECMHHE",
	AGOKLPNCBDN:                               "AGOKLPNCBDN",
	BKHONFCFGKB:                               "BKHONFCFGKB",
	LIHDIHMOLGB:                               "LIHDIHMOLGB",
	HHOHOHHFJLF:                               "HHOHOHHFJLF",
	KMLLPPEIBOP:                               "KMLLPPEIBOP",
	EraFlipperDataChangeScNotify:              "EraFlipperDataChangeScNotify",
	MBJFFJDBAJJ:                               "MBJFFJDBAJJ",
	GDPEMIMGBNI:                               "GDPEMIMGBNI",
	CLBIFDFPJNK:                               "CLBIFDFPJNK",
	HOMKHAKNBCD:                               "HOMKHAKNBCD",
	DCFMOCENKKJ:                               "DCFMOCENKKJ",
	HMDKLFGHGMG:                               "HMDKLFGHGMG",
	ChangeEraFlipperDataCsReq:                 "ChangeEraFlipperDataCsReq",
	ResetEraFlipperDataScRsp:                  "ResetEraFlipperDataScRsp",
	FKCIBEMNKFE:                               "FKCIBEMNKFE",
	EnterEraFlipperRegionScRsp:                "EnterEraFlipperRegionScRsp",
	OLPAAHBAHMJ:                               "OLPAAHBAHMJ",
	CNLBFFIELGM:                               "CNLBFFIELGM",
	ADEBBLIHHNO:                               "ADEBBLIHHNO",
	FPCCJFBKNFI:                               "FPCCJFBKNFI",
	FDECMDAKEDC:                               "FDECMDAKEDC",
	DOLBDKNMMEL:                               "DOLBDKNMMEL",
	GPHFMAKAMPG:                               "GPHFMAKAMPG",
	GPBIAGAOKKB:                               "GPBIAGAOKKB",
	KCNPMHJGHNM:                               "KCNPMHJGHNM",
	FJMDLCBMMFJ:                               "FJMDLCBMMFJ",
	ClockParkGetInfoScRsp:                     "ClockParkGetInfoScRsp",
	BPMLEFKKJCG:                               "BPMLEFKKJCG",
	FEIPJNNPMOC:                               "FEIPJNNPMOC",
	JLMEMDMFKGM:                               "JLMEMDMFKGM",
	BMDPJKOGJDM:                               "BMDPJKOGJDM",
	NNKEMLDCGGP:                               "NNKEMLDCGGP",
	LFODGNFBPOB:                               "LFODGNFBPOB",
	ClockParkGetInfoCsReq:                     "ClockParkGetInfoCsReq",
	FKNHMACBHLP:                               "FKNHMACBHLP",
	CCHIMMEMKBO:                               "CCHIMMEMKBO",
	JIOBHLOPKHB:                               "JIOBHLOPKHB",
	FAGPOLDODGP:                               "FAGPOLDODGP",
	HCBIODPEOHE:                               "HCBIODPEOHE",
	EBGILDGNGCL:                               "EBGILDGNGCL",
	JMCMMNKANEO:                               "JMCMMNKANEO",
	HPDKMOJIGNE:                               "HPDKMOJIGNE",
	DOEAMNCKPNN:                               "DOEAMNCKPNN",
	CGAAGOKHIEM:                               "CGAAGOKHIEM",
	KFAGMGMMHGC:                               "KFAGMGMMHGC",
	CJDINMHAICJ:                               "CJDINMHAICJ",
	GKIALGMGJND:                               "GKIALGMGJND",
	GBFLDCLCEDO:                               "GBFLDCLCEDO",
	LMAOLBGBFHH:                               "LMAOLBGBFHH",
	HCKPAANCJEH:                               "HCKPAANCJEH",
	StartMatchScRsp:                           "StartMatchScRsp",
	CancelMatchScRsp:                          "CancelMatchScRsp",
	GetCrossInfoScRsp:                         "GetCrossInfoScRsp",
	GetCrossInfoCsReq:                         "GetCrossInfoCsReq",
	MatchResultScNotify:                       "MatchResultScNotify",
	StartMatchCsReq:                           "StartMatchCsReq",
	EMIOJOLOCPM:                               "EMIOJOLOCPM",
	IGKLAOPJFOG:                               "IGKLAOPJFOG",
	GEPPCEOCLOG:                               "GEPPCEOCLOG",
	PIHIECEMCCB:                               "PIHIECEMCCB",
	DMHLEHMCMJA:                               "DMHLEHMCMJA",
	HIIENOHBGPG:                               "HIIENOHBGPG",
	HDKFBPGAOAM:                               "HDKFBPGAOAM",
	OCHFADEBKHA:                               "OCHFADEBKHA",
	ELOCPGINHJC:                               "ELOCPGINHJC",
	FBIALGFBAKH:                               "FBIALGFBAKH",
	IFAPOJIDINH:                               "IFAPOJIDINH",
	DBLHNDBNJHF:                               "DBLHNDBNJHF",
	EOGBKAIKCKE:                               "EOGBKAIKCKE",
	NNIEBHKIIIP:                               "NNIEBHKIIIP",
	GOGOMFCOHJA:                               "GOGOMFCOHJA",
	HINLJEPMAEO:                               "HINLJEPMAEO",
	EKHNLDEELLH:                               "EKHNLDEELLH",
	KDHMNKGIADL:                               "KDHMNKGIADL",
	AENCOEBKFMA:                               "AENCOEBKFMA",
	NMIOCHECKLE:                               "NMIOCHECKLE",
	DAADKJADMPP:                               "DAADKJADMPP",
	NDNODNGCIFP:                               "NDNODNGCIFP",
	ANFGNJFOFFI:                               "ANFGNJFOFFI",
	JFHBMCMGGPG:                               "JFHBMCMGGPG",
	BEHGLJAPOJF:                               "BEHGLJAPOJF",
	HEOOEKCDMBA:                               "HEOOEKCDMBA",
	KFPNNLAGCON:                               "KFPNNLAGCON",
	OLIPBINKPHL:                               "OLIPBINKPHL",
	IPBNIMPLEBD:                               "IPBNIMPLEBD",
	GOFPCLNNHLG:                               "GOFPCLNNHLG",
	AGLLFHMLCFF:                               "AGLLFHMLCFF",
	BNPJMLIIJHE:                               "BNPJMLIIJHE",
	EMKJJGADDLD:                               "EMKJJGADDLD",
	EEMDHJFINBD:                               "EEMDHJFINBD",
	LNBENIECPOI:                               "LNBENIECPOI",
	BODCPLJDPGF:                               "BODCPLJDPGF",
	PENCGFNCOEI:                               "PENCGFNCOEI",
	BBGHLIMPOBH:                               "BBGHLIMPOBH",
	LCNHIJFMBAI:                               "LCNHIJFMBAI",
	NJKMOLIKBAK:                               "NJKMOLIKBAK",
	CHFOOHMJNPO:                               "CHFOOHMJNPO",
	HNMPDDELBMP:                               "HNMPDDELBMP",
	CKIAKLCJJLH:                               "CKIAKLCJJLH",
	JNPHKPDLOIG:                               "JNPHKPDLOIG",
	GGLAHPHLDCO:                               "GGLAHPHLDCO",
	NNEDJEHPIIE:                               "NNEDJEHPIIE",
	AJBANMKILCG:                               "AJBANMKILCG",
	EGCMACJOKME:                               "EGCMACJOKME",
	EMALGPDKFAD:                               "EMALGPDKFAD",
	JALNCLCCNJG:                               "JALNCLCCNJG",
	LOKDBDKFNJF:                               "LOKDBDKFNJF",
	JIFBPGHCAFF:                               "JIFBPGHCAFF",
	KPLBPAOOCJI:                               "KPLBPAOOCJI",
	ANMLGADENCM:                               "ANMLGADENCM",
	MHJDEDOPLLJ:                               "MHJDEDOPLLJ",
	JKMPIABFGIO:                               "JKMPIABFGIO",
	BFGJOPMMELH:                               "BFGJOPMMELH",
	EMFHGILAAMN:                               "EMFHGILAAMN",
	CMPCDMFEEBH:                               "CMPCDMFEEBH",
	IMFEIOIEJMP:                               "IMFEIOIEJMP",
	PCGEBNILLME:                               "PCGEBNILLME",
	JOMHCPMBFBB:                               "JOMHCPMBFBB",
	IPONIEACFBD:                               "IPONIEACFBD",
	CKLNHPPPAHM:                               "CKLNHPPPAHM",
	CPJDLAKBGAF:                               "CPJDLAKBGAF",
	DFCHACFIOCP:                               "DFCHACFIOCP",
	EGPNAGDIMBK:                               "EGPNAGDIMBK",
	ELACOJOEBMC:                               "ELACOJOEBMC",
	JLADNFEAAEI:                               "JLADNFEAAEI",
	PABCOMDNKOG:                               "PABCOMDNKOG",
	MBMAOKAKEPA:                               "MBMAOKAKEPA",
	MEGEDCOBMKC:                               "MEGEDCOBMKC",
	BCAFEPLNPGG:                               "BCAFEPLNPGG",
	KPIKJKNMFKE:                               "KPIKJKNMFKE",
	PGGINNBANBD:                               "PGGINNBANBD",
	ContentPackageGetDataScRsp:                "ContentPackageGetDataScRsp",
	ContentPackageGetDataCsReq:                "ContentPackageGetDataCsReq",
	ContentPackageTransferScNotify:            "ContentPackageTransferScNotify",
	ContentPackageSyncDataScNotify:            "ContentPackageSyncDataScNotify",
	MIHHADFCMLC:                               "MIHHADFCMLC",
	FDNIEBAOINP:                               "FDNIEBAOINP",
	DFDDNHAMHHC:                               "DFDDNHAMHHC",
	IDIEFCGCKBD:                               "IDIEFCGCKBD",
	AHAMDHHLPEN:                               "AHAMDHHLPEN",
	FPDOEGNIAAK:                               "FPDOEGNIAAK",
	AFBBLGEHOBE:                               "AFBBLGEHOBE",
	OKHPABFGLOG:                               "OKHPABFGLOG",
	GHEJNPKMHCP:                               "GHEJNPKMHCP",
	GEKLGPLOBOE:                               "GEKLGPLOBOE",
	BPIPPBJBFBE:                               "BPIPPBJBFBE",
	NFKANNMMPEH:                               "NFKANNMMPEH",
	ICIBMIJKHAK:                               "ICIBMIJKHAK",
	MusicRhythmStartLevelCsReq:                "MusicRhythmStartLevelCsReq",
	MusicRhythmFinishLevelCsReq:               "MusicRhythmFinishLevelCsReq",
	MusicRhythmMaxDifficultyLevelsUnlockNotify: "MusicRhythmMaxDifficultyLevelsUnlockNotify",
	MusicRhythmDataScRsp:                       "MusicRhythmDataScRsp",
	MusicRhythmUnlockSongNotify:                "MusicRhythmUnlockSongNotify",
	MusicRhythmStartLevelScRsp:                 "MusicRhythmStartLevelScRsp",
	MusicRhythmFinishLevelScRsp:                "MusicRhythmFinishLevelScRsp",
	MusicRhythmSaveSongConfigDataScRsp:         "MusicRhythmSaveSongConfigDataScRsp",
	MusicRhythmUnlockTrackScNotify:             "MusicRhythmUnlockTrackScNotify",
	MusicRhythmDataCsReq:                       "MusicRhythmDataCsReq",
	MusicRhythmUnlockSongSfxScNotify:           "MusicRhythmUnlockSongSfxScNotify",
	MusicRhythmSaveSongConfigDataCsReq:         "MusicRhythmSaveSongConfigDataCsReq",
	CurPetChangedScNotify:                      "CurPetChangedScNotify",
	GetPetDataScRsp:                            "GetPetDataScRsp",
	SummonPetScRsp:                             "SummonPetScRsp",
	GetPetDataCsReq:                            "GetPetDataCsReq",
	RecallPetCsReq:                             "RecallPetCsReq",
	RecallPetScRsp:                             "RecallPetScRsp",
	SummonPetCsReq:                             "SummonPetCsReq",
	WorldUnlockScRsp:                           "WorldUnlockScRsp",
	WorldUnlockCsReq:                           "WorldUnlockCsReq",
	LCAMCNOFPMP:                                "LCAMCNOFPMP",
	LMAHEGLKOJJ:                                "LMAHEGLKOJJ",
	PKKLLLANPFG:                                "PKKLLLANPFG",
	PGFGEEADLJP:                                "PGFGEEADLJP",
	FKDEBAICAAO:                                "FKDEBAICAAO",
	PBNFEGELHPB:                                "PBNFEGELHPB",
	BAEODIMMFGE:                                "BAEODIMMFGE",
	KGCHIHCEINK:                                "KGCHIHCEINK",
	CLOAAFAPKPG:                                "CLOAAFAPKPG",
	IGHONOMDLHL:                                "IGHONOMDLHL",
	PLPOBHLIKLC:                                "PLPOBHLIKLC",
	LLAAGPNGDAI:                                "LLAAGPNGDAI",
	PBCFOGEGLMB:                                "PBCFOGEGLMB",
	BBNOJFFAPCP:                                "BBNOJFFAPCP",
	MNEBHGMAGGD:                                "MNEBHGMAGGD",
	APEGFHBBGDD:                                "APEGFHBBGDD",
	DMPAINJEGCD:                                "DMPAINJEGCD",
	DFAKLHGPBCJ:                                "DFAKLHGPBCJ",
	DMKHEJINBHI:                                "DMKHEJINBHI",
	PAOFAEHLGEJ:                                "PAOFAEHLGEJ",
	ALMEAJGADBP:                                "ALMEAJGADBP",
	JDOFCFELEMH:                                "JDOFCFELEMH",
	LHGOAKJKPDF:                                "LHGOAKJKPDF",
	FBKGCCFALGG:                                "FBKGCCFALGG",
	GOILEIMINDK:                                "GOILEIMINDK",
	NKPDDCNPJPE:                                "NKPDDCNPJPE",
	BAFECPJFFJJ:                                "BAFECPJFFJJ",
	OFMJLFKGBEA:                                "OFMJLFKGBEA",
	GLOPKOBPOIB:                                "GLOPKOBPOIB",
	LAMFNNNCNNL:                                "LAMFNNNCNNL",
	AMEJIHNHFID:                                "AMEJIHNHFID",
	EKGPBLKOCFH:                                "EKGPBLKOCFH",
	OKHKBMNLLIL:                                "OKHKBMNLLIL",
	FBAIJOOPECG:                                "FBAIJOOPECG",
	OKFNCCGBHHI:                                "OKFNCCGBHHI",
	AMNPFHHPMGB:                                "AMNPFHHPMGB",
	DPOEEJJBLEC:                                "DPOEEJJBLEC",
	POFMBILCPMK:                                "POFMBILCPMK",
	FKEOCMNDHAN:                                "FKEOCMNDHAN",
	CFGHHOMOFJF:                                "CFGHHOMOFJF",
	GEDKMFICABA:                                "GEDKMFICABA",
	PNJNIGHFNDA:                                "PNJNIGHFNDA",
	FFBNNEMNAIG:                                "FFBNNEMNAIG",
	GIFKNBFOPDP:                                "GIFKNBFOPDP",
	FOGJMOKNDCL:                                "FOGJMOKNDCL",
	CNNONAOENKL:                                "CNNONAOENKL",
	AKNPOFCGAEA:                                "AKNPOFCGAEA",
	TrainPartyGetDataScRsp:                     "TrainPartyGetDataScRsp",
	TrainPartyEnterCsReq:                       "TrainPartyEnterCsReq",
	TrainPartyGamePlaySettleNotify:             "TrainPartyGamePlaySettleNotify",
	TrainPartyUseCardScRsp:                     "TrainPartyUseCardScRsp",
	MFFCOHOKBHD:                                "MFFCOHOKBHD",
	HANHHJNHLEM:                                "HANHHJNHLEM",
	BPAHALHPHJB:                                "BPAHALHPHJB",
	TrainPartyLeaveScRsp:                       "TrainPartyLeaveScRsp",
	TrainPartyUseCardCsReq:                     "TrainPartyUseCardCsReq",
	TrainPartyHandlePendingActionScRsp:         "TrainPartyHandlePendingActionScRsp",
	TrainPartyGetDataCsReq:                     "TrainPartyGetDataCsReq",
	IFPJKDDMGEP:                                "IFPJKDDMGEP",
	TrainPartyBuildDiyCsReq:                    "TrainPartyBuildDiyCsReq",
	TrainPartyMoveScNotify:                     "TrainPartyMoveScNotify",
	TrainPartyBuildStartStepScRsp:              "TrainPartyBuildStartStepScRsp",
	MOCCJFOHAAB:                                "MOCCJFOHAAB",
	LKOGDOIABBE:                                "LKOGDOIABBE",
	TrainPartyHandlePendingActionCsReq:         "TrainPartyHandlePendingActionCsReq",
	TrainPartyLeaveCsReq:                       "TrainPartyLeaveCsReq",
	TrainPartyEnterScRsp:                       "TrainPartyEnterScRsp",
	TrainPartyBuildRoomScNotify:                "TrainPartyBuildRoomScNotify",
	TrainPartyBuildStartStepCsReq:              "TrainPartyBuildStartStepCsReq",
	MJNLNBFAOGA:                                "MJNLNBFAOGA",
	TrainPartyBuildDiyScRsp:                    "TrainPartyBuildDiyScRsp",
	TrainPartySettleNotify:                     "TrainPartySettleNotify",
	TrainPartyBuildingUpdateNotify:             "TrainPartyBuildingUpdateNotify",
	TrainPartySyncUpdateScNotify:               "TrainPartySyncUpdateScNotify",
	EGNAJNKNGGG:                                "EGNAJNKNGGG",
	HECMLPGBPEP:                                "HECMLPGBPEP",
	EJPMLFBNIJL:                                "EJPMLFBNIJL",
	NAJNNIJNAHO:                                "NAJNNIJNAHO",
	SwitchHandDataScRsp:                        "SwitchHandDataScRsp",
	JDNLHLGKABA:                                "JDNLHLGKABA",
	GDLIDFPKEFN:                                "GDLIDFPKEFN",
	APCOIFHMLGJ:                                "APCOIFHMLGJ",
	AGKHKDAEBLL:                                "AGKHKDAEBLL",
	SwitchHandDataCsReq:                        "SwitchHandDataCsReq",
	GCPCNFLCLJJ:                                "GCPCNFLCLJJ",
	DEKBEKIDNKL:                                "DEKBEKIDNKL",
	MDBHIKBGFAE:                                "MDBHIKBGFAE",
	DAKJIBFCIKD:                                "DAKJIBFCIKD",
	SelectPamSkinScRsp:                         "SelectPamSkinScRsp",
	GetPamSkinDataScRsp:                        "GetPamSkinDataScRsp",
	SelectPamSkinCsReq:                         "SelectPamSkinCsReq",
	GetPamSkinDataCsReq:                        "GetPamSkinDataCsReq",
	UnlockPamSkinScNotify:                      "UnlockPamSkinScNotify",
	IKALGCPBNDP:                                "IKALGCPBNDP",
	LODHDFDJIOB:                                "LODHDFDJIOB",
	ABKDDBIDCJE:                                "ABKDDBIDCJE",
	NMNNAGFIPEP:                                "NMNNAGFIPEP",
	GJOCPAMNIIH:                                "GJOCPAMNIIH",
	JMKEBFCMIPO:                                "JMKEBFCMIPO",
	KINPGIFGGED:                                "KINPGIFGGED",
	KMENEDCHAAM:                                "KMENEDCHAAM",
	HHKCFNGDHKL:                                "HHKCFNGDHKL",
	EAAPHCNFLEK:                                "EAAPHCNFLEK",
	LMHFDBJPCAP:                                "LMHFDBJPCAP",
	BMKBLIHIKIG:                                "BMKBLIHIKIG",
	HKMIMDPJAJN:                                "HKMIMDPJAJN",
	DFELMHMHDOP:                                "DFELMHMHDOP",
	NKLPPGBOOJN:                                "NKLPPGBOOJN",
	FPMPFCJIOCP:                                "FPMPFCJIOCP",
	JIMMKLIHONA:                                "JIMMKLIHONA",
	MMPAMLGKGFG:                                "MMPAMLGKGFG",
	ICLGFIHJIAF:                                "ICLGFIHJIAF",
	EABDAHBDKPE:                                "EABDAHBDKPE",
	HBEDPFHCAPJ:                                "HBEDPFHCAPJ",
	BOJIMHGFGEE:                                "BOJIMHGFGEE",
	EBNIKJNJMBD:                                "EBNIKJNJMBD",
	IMEIPJBFLPL:                                "IMEIPJBFLPL",
	OKOMOBEDNPF:                                "OKOMOBEDNPF",
	IOMKBPMHPAM:                                "IOMKBPMHPAM",
	HKFIBMBCPNK:                                "HKFIBMBCPNK",
	UpdateMarkChestScRsp:                       "UpdateMarkChestScRsp",
	GetMarkChestScRsp:                          "GetMarkChestScRsp",
	UpdateMarkChestCsReq:                       "UpdateMarkChestCsReq",
	GetMarkChestCsReq:                          "GetMarkChestCsReq",
	MarkChestChangedScNotify:                   "MarkChestChangedScNotify",
	EHLDJKJGNCN:                                "EHLDJKJGNCN",
	CHOIPPBOANF:                                "CHOIPPBOANF",
	OGHKDPJINAK:                                "OGHKDPJINAK",
	HOJLKBMHNPH:                                "HOJLKBMHNPH",
	DNDEJBMDCMJ:                                "DNDEJBMDCMJ",
	DJLLAEPEBHE:                                "DJLLAEPEBHE",
	MPKEEIMNGHI:                                "MPKEEIMNGHI",
	DJBMCJGOLBC:                                "DJBMCJGOLBC",
	FMLCDNJJLMC:                                "FMLCDNJJLMC",
	JBHLHNKJKOF:                                "JBHLHNKJKOF",
	KFDEIFMMHIC:                                "KFDEIFMMHIC",
	FINKJPOOIKG:                                "FINKJPOOIKG",
	MJOINJMAHBI:                                "MJOINJMAHBI",
	HFDHGMDBEJF:                                "HFDHGMDBEJF",
	FDBDPJJPDHO:                                "FDBDPJJPDHO",
	FCECPHKCHGK:                                "FCECPHKCHGK",
	JFJGLMKBAFD:                                "JFJGLMKBAFD",
	OGKFHINGNLC:                                "OGKFHINGNLC",
	LGNMIJKJKJC:                                "LGNMIJKJKJC",
	DKEOICABNFP:                                "DKEOICABNFP",
	HFAJLDKAMEO:                                "HFAJLDKAMEO",
	NDCAPHOBAKC:                                "NDCAPHOBAKC",
	BNOILJLCIOE:                                "BNOILJLCIOE",
	MIFINMNPMBP:                                "MIFINMNPMBP",
	CIFAHBDLDED:                                "CIFAHBDLDED",
	KKMKKBDBMNB:                                "KKMKKBDBMNB",
	AAIANPMCIBG:                                "AAIANPMCIBG",
	HMKNNAJKAGE:                                "HMKNNAJKAGE",
	DKIJNKLPBFF:                                "DKIJNKLPBFF",
	FOOLDOEIFOC:                                "FOOLDOEIFOC",
	DANCJEECPMN:                                "DANCJEECPMN",
	LPAMFAMAFEP:                                "LPAMFAMAFEP",
	BCPCENAKCJK:                                "BCPCENAKCJK",
	PFHHIGHEDJK:                                "PFHHIGHEDJK",
	AAJFBJHMAON:                                "AAJFBJHMAON",
	LNOLPAEGLGI:                                "LNOLPAEGLGI",
	HDMEHMJLNAB:                                "HDMEHMJLNAB",
	HEBMHDINMIP:                                "HEBMHDINMIP",
	LKLAMDJOPCE:                                "LKLAMDJOPCE",
	DOCFADPCGIO:                                "DOCFADPCGIO",
	OKPIJHJGAFJ:                                "OKPIJHJGAFJ",
	KLILCFENABP:                                "KLILCFENABP",
	KIHDMHNCAAL:                                "KIHDMHNCAAL",
	DGKOFIJEHLH:                                "DGKOFIJEHLH",
	OLBOENJMCNL:                                "OLBOENJMCNL",
	JAGIJPGIBHF:                                "JAGIJPGIBHF",
	HEEDAPEOIHG:                                "HEEDAPEOIHG",
	POKLKNJFJJP:                                "POKLKNJFJJP",
	OLHNBGKLKDL:                                "OLHNBGKLKDL",
	PLBFHOPECGI:                                "PLBFHOPECGI",
	RelicSmartWearUpdatePinRelicCsReq:          "RelicSmartWearUpdatePinRelicCsReq",
	RelicSmartWearDeletePlanCsReq:              "RelicSmartWearDeletePlanCsReq",
	RelicSmartWearUpdatePlanScRsp:              "RelicSmartWearUpdatePlanScRsp",
	RelicSmartWearAddPlanScRsp:                 "RelicSmartWearAddPlanScRsp",
	RelicSmartWearGetPlanScRsp:                 "RelicSmartWearGetPlanScRsp",
	RelicSmartWearUpdatePinRelicScNotify:       "RelicSmartWearUpdatePinRelicScNotify",
	RelicSmartWearAddPlanCsReq:                 "RelicSmartWearAddPlanCsReq",
	RelicSmartWearGetPinRelicCsReq:             "RelicSmartWearGetPinRelicCsReq",
	RelicSmartWearGetPlanCsReq:                 "RelicSmartWearGetPlanCsReq",
	RelicSmartWearUpdatePlanCsReq:              "RelicSmartWearUpdatePlanCsReq",
	RelicSmartWearDeletePinRelicScRsp:          "RelicSmartWearDeletePinRelicScRsp",
	RelicSmartWearDeletePlanScRsp:              "RelicSmartWearDeletePlanScRsp",
	RelicSmartWearUpdatePinRelicScRsp:          "RelicSmartWearUpdatePinRelicScRsp",
	RelicSmartWearGetPinRelicScRsp:             "RelicSmartWearGetPinRelicScRsp",
	OLJEAPHMGIE:                                "OLJEAPHMGIE",
	GBHLODAEAHH:                                "GBHLODAEAHH",
	LLBMDLJHJLO:                                "LLBMDLJHJLO",
	BPNJPCNEKBD:                                "BPNJPCNEKBD",
	EGHOBCIMPNH:                                "EGHOBCIMPNH",
	MFANAADHHPG:                                "MFANAADHHPG",
	PNIEHBLNHOB:                                "PNIEHBLNHOB",
	KGEIJDNMEKO:                                "KGEIJDNMEKO",
	KCGJFHILEOI:                                "KCGJFHILEOI",
	OKEKJDFOIKD:                                "OKEKJDFOIKD",
	IAGEJNPIHJG:                                "IAGEJNPIHJG",
	CAOGKPDKPID:                                "CAOGKPDKPID",
	BDFNNPAPBOJ:                                "BDFNNPAPBOJ",
	GEGJIOCOOBI:                                "GEGJIOCOOBI",
	FJBFEBJNBHK:                                "FJBFEBJNBHK",
	EJOBJPJIEMF:                                "EJOBJPJIEMF",
	FIAMCGLOJDE:                                "FIAMCGLOJDE",
	EOPBCNGDODJ:                                "EOPBCNGDODJ",
	FNDNKAACBHD:                                "FNDNKAACBHD",
	HCBGBMKMAAA:                                "HCBGBMKMAAA",
	BJOCIHICCCM:                                "BJOCIHICCCM",
	IPJEPOGIDFL:                                "IPJEPOGIDFL",
	CLIPNNJBNFI:                                "CLIPNNJBNFI",
	HMAFHKAEKFG:                                "HMAFHKAEKFG",
	JLHCAEHNBEJ:                                "JLHCAEHNBEJ",
	EBKEJLJCMOP:                                "EBKEJLJCMOP",
	KNCIJNEPLHN:                                "KNCIJNEPLHN",
	JCBDEKIGMGC:                                "JCBDEKIGMGC",
	SyncRechargeBenefitInfoScNotify:            "SyncRechargeBenefitInfoScNotify",
	GetRechargeBenefitInfoScRsp:                "GetRechargeBenefitInfoScRsp",
	TakeRechargeGiftRewardScRsp:                "TakeRechargeGiftRewardScRsp",
	GetRechargeGiftInfoScRsp:                   "GetRechargeGiftInfoScRsp",
	TakeRechargeGiftRewardCsReq:                "TakeRechargeGiftRewardCsReq",
	TakeRechargeBenefitRewardScRsp:             "TakeRechargeBenefitRewardScRsp",
	GetRechargeGiftInfoCsReq:                   "GetRechargeGiftInfoCsReq",
	GetRechargeBenefitInfoCsReq:                "GetRechargeBenefitInfoCsReq",
	TakeRechargeBenefitRewardCsReq:             "TakeRechargeBenefitRewardCsReq",
	LLIEMGCGHDG:                                "LLIEMGCGHDG",
	LPIKOGPOMNI:                                "LPIKOGPOMNI",
	PPMGKFEDNOL:                                "PPMGKFEDNOL",
	GNLLJMIMNPC:                                "GNLLJMIMNPC",
	FPMFIIPPIMO:                                "FPMFIIPPIMO",
	HDBILLENKMG:                                "HDBILLENKMG",
	KIFFLDHFMLL:                                "KIFFLDHFMLL",
	BHHDIGKNGGK:                                "BHHDIGKNGGK",
	GHBGNOKPDKL:                                "GHBGNOKPDKL",
	IKOMIMHGIIN:                                "IKOMIMHGIIN",
	JCMOOEBEONF:                                "JCMOOEBEONF",
	FEMLPCAFCFE:                                "FEMLPCAFCFE",
	KPBGJHIHKHA:                                "KPBGJHIHKHA",
	EPMNHLAEJOD:                                "EPMNHLAEJOD",
	FIGGHKDLDEH:                                "FIGGHKDLDEH",
	KHNMPOADPFM:                                "KHNMPOADPFM",
	DBNEGFOMOJG:                                "DBNEGFOMOJG",
	BEOPBDOOPPD:                                "BEOPBDOOPPD",
	BHFCGFIIPEK:                                "BHFCGFIIPEK",
	IEHCLHCBIAP:                                "IEHCLHCBIAP",
	BKLFEDBGIOP:                                "BKLFEDBGIOP",
	KMEOJGKPNIC:                                "KMEOJGKPNIC",
	GPIBEDECNDP:                                "GPIBEDECNDP",
	IJNCIDCCNFG:                                "IJNCIDCCNFG",
	CJHKHIJCBDL:                                "CJHKHIJCBDL",
	KPMLKCDIFBE:                                "KPMLKCDIFBE",
	JOCPDPNOBNA:                                "JOCPDPNOBNA",
	DLKHMMLJHCA:                                "DLKHMMLJHCA",
	FJDAGBIJMEP:                                "FJDAGBIJMEP",
	EKOOANIDBPP:                                "EKOOANIDBPP",
	PGHHJGLLOKP:                                "PGHHJGLLOKP",
	FCMPHKJEBJA:                                "FCMPHKJEBJA",
	NCGOHBCBFMH:                                "NCGOHBCBFMH",
	EHBEPEFDPHC:                                "EHBEPEFDPHC",
	PBPIGABCJED:                                "PBPIGABCJED",
	EKELNGFIDJD:                                "EKELNGFIDJD",
	NFPBIDKBCLM:                                "NFPBIDKBCLM",
	CNOINGLLHMB:                                "CNOINGLLHMB",
	MMJAGLCLIPO:                                "MMJAGLCLIPO",
	GOIIALEIBOB:                                "GOIIALEIBOB",
	HEHOBIPOHCP:                                "HEHOBIPOHCP",
	MIFGNJIMLKO:                                "MIFGNJIMLKO",
	GLFKBPCDNLM:                                "GLFKBPCDNLM",
	AHBFFEPLBMD:                                "AHBFFEPLBMD",
	CBPGKBBBEEP:                                "CBPGKBBBEEP",
	HHGHILJDNAI:                                "HHGHILJDNAI",
	GDBEHBFBKKD:                                "GDBEHBFBKKD",
	MMJBHOMPNNH:                                "MMJBHOMPNNH",
	NDLGFKKELLN:                                "NDLGFKKELLN",
	KNGCCCKIGDF:                                "KNGCCCKIGDF",
	GFDGPBELEGB:                                "GFDGPBELEGB",
	DLKCLHJNEGB:                                "DLKCLHJNEGB",
	KFHNBIKNLBH:                                "KFHNBIKNLBH",
	KMKOFIKGLJD:                                "KMKOFIKGLJD",
	OKJBBHPFFAN:                                "OKJBBHPFFAN",
	HJFBJFNPDKB:                                "HJFBJFNPDKB",
	JDJFNMFEHNG:                                "JDJFNMFEHNG",
	ABJBJOCBPLH:                                "ABJBJOCBPLH",
	CGDADAACBFO:                                "CGDADAACBFO",
	FCDPNMDNGFC:                                "FCDPNMDNGFC",
	NENOCOGHFON:                                "NENOCOGHFON",
	DMKLPOKPIGE:                                "DMKLPOKPIGE",
	FENMMJBPKKI:                                "FENMMJBPKKI",
	FAGAFKIDIKB:                                "FAGAFKIDIKB",
	LEDECKHLPJN:                                "LEDECKHLPJN",
	GMCAIKPAGMC:                                "GMCAIKPAGMC",
	GKILPAEELJH:                                "GKILPAEELJH",
	KACCPALJOFK:                                "KACCPALJOFK",
	NDAJHDMMCOI:                                "NDAJHDMMCOI",
	IBOBBDILLEF:                                "IBOBBDILLEF",
	NBKMJNMLBBO:                                "NBKMJNMLBBO",
	NPFMIPDCLOI:                                "NPFMIPDCLOI",
	HEHLILCFMME:                                "HEHLILCFMME",
	GMAAJCPDFBF:                                "GMAAJCPDFBF",
	BNANEKNNMGD:                                "BNANEKNNMGD",
	DELMBJCMFAP:                                "DELMBJCMFAP",
	LCNHOJPJKHM:                                "LCNHOJPJKHM",
	PNNJIKHKIHJ:                                "PNNJIKHKIHJ",
	MGCCIGAGECD:                                "MGCCIGAGECD",
	DOEKJBOHGGD:                                "DOEKJBOHGGD",
	NOFKLGLKCAL:                                "NOFKLGLKCAL",
	BANKFKMLPKP:                                "BANKFKMLPKP",
	OFHPHMDAEAB:                                "OFHPHMDAEAB",
	PGKDCDONKEI:                                "PGKDCDONKEI",
	NDHPMCECJDN:                                "NDHPMCECJDN",
	EBOAHJGEKMP:                                "EBOAHJGEKMP",
	ONBJLECGIID:                                "ONBJLECGIID",
	GetChallengePeakDataScRsp:                  "GetChallengePeakDataScRsp",
	StartChallengePeakScRsp:                    "StartChallengePeakScRsp",
	GetCurChallengePeakScRsp:                   "GetCurChallengePeakScRsp",
	SetChallengePeakMobLineupAvatarCsReq:       "SetChallengePeakMobLineupAvatarCsReq",
	GetCurChallengePeakCsReq:                   "GetCurChallengePeakCsReq",
	SetChallengePeakBossHardModeCsReq:          "SetChallengePeakBossHardModeCsReq",
	ConfirmChallengePeakSettleScRsp:            "ConfirmChallengePeakSettleScRsp",
	TakeChallengePeakRewardScRsp:               "TakeChallengePeakRewardScRsp",
	GetChallengePeakDataCsReq:                  "GetChallengePeakDataCsReq",
	ReStartChallengePeakScRsp:                  "ReStartChallengePeakScRsp",
	ConfirmChallengePeakSettleCsReq:            "ConfirmChallengePeakSettleCsReq",
	ChallengePeakGroupDataUpdateScNotify:       "ChallengePeakGroupDataUpdateScNotify",
	ChallengePeakSettleScNotify:                "ChallengePeakSettleScNotify",
	LeaveChallengePeakScRsp:                    "LeaveChallengePeakScRsp",
	ReStartChallengePeakCsReq:                  "ReStartChallengePeakCsReq",
	LeaveChallengePeakCsReq:                    "LeaveChallengePeakCsReq",
	SetChallengePeakBossHardModeScRsp:          "SetChallengePeakBossHardModeScRsp",
	TakeChallengePeakRewardCsReq:               "TakeChallengePeakRewardCsReq",
	StartChallengePeakCsReq:                    "StartChallengePeakCsReq",
	SetChallengePeakMobLineupAvatarScRsp:       "SetChallengePeakMobLineupAvatarScRsp",
	FKBJJIMKFOK:                                "FKBJJIMKFOK",
	NPGONPEDINB:                                "NPGONPEDINB",
	IDIKMHHOIIC:                                "IDIKMHHOIIC",
	NJHPBNBHOLL:                                "NJHPBNBHOLL",
	CAICLJOILEK:                                "CAICLJOILEK",
	GJMDFCALIOL:                                "GJMDFCALIOL",
	NDINECKDMDO:                                "NDINECKDMDO",
	EKNCLPMCCNE:                                "EKNCLPMCCNE",
	JNKIOGAEJDG:                                "JNKIOGAEJDG",
	FPLOMPFOFPL:                                "FPLOMPFOFPL",
	JIAMLEMCLHD:                                "JIAMLEMCLHD",
	GOHAIMEMEFG:                                "GOHAIMEMEFG",
	MLGECIDINHM:                                "MLGECIDINHM",
	LKJMMAJJCBK:                                "LKJMMAJJCBK",
	IMKPBFNLMBJ:                                "IMKPBFNLMBJ",
	MCHDFOHHBLF:                                "MCHDFOHHBLF",
	PKLFFNMIIAN:                                "PKLFFNMIIAN",
	JBFCBGJADOA:                                "JBFCBGJADOA",
	IJJLBEKHICO:                                "IJJLBEKHICO",
	OHFABPDBIDE:                                "OHFABPDBIDE",
	PFCGHODONEK:                                "PFCGHODONEK",
	FBPHNLFDFAE:                                "FBPHNLFDFAE",
	HJJAFGIBNDI:                                "HJJAFGIBNDI",
	GIGAMCJGJKP:                                "GIGAMCJGJKP",
	AJDBHHLKLBB:                                "AJDBHHLKLBB",
	GMJIHGDONDE:                                "GMJIHGDONDE",
	MHOAJAJNBGH:                                "MHOAJAJNBGH",
	JBGBOMDDOEN:                                "JBGBOMDDOEN",
	JIHBNKNCJKH:                                "JIHBNKNCJKH",
	ANPNFCFDDLA:                                "ANPNFCFDDLA",
	EnterElationActivityStageScRsp:             "EnterElationActivityStageScRsp",
	GetElationActivityDataCsReq:                "GetElationActivityDataCsReq",
	ElationActivityBattleEndScNotify:           "ElationActivityBattleEndScNotify",
	GetElationActivityDataScRsp:                "GetElationActivityDataScRsp",
	EnterElationActivityStageCsReq:             "EnterElationActivityStageCsReq",
	GetActivityRewardCountDataScRsp:            "GetActivityRewardCountDataScRsp",
	GetActivityHotDataCsReq:                    "GetActivityHotDataCsReq",
	GetActivityHotDataScRsp:                    "GetActivityHotDataScRsp",
	GetActivityRewardCountDataCsReq:            "GetActivityRewardCountDataCsReq",
	FHPKLLOPKOC:                                "FHPKLLOPKOC",
	JEFONNOLAEB:                                "JEFONNOLAEB",
	KEMCCEIJAKD:                                "KEMCCEIJAKD",
	EHGALIKLNHA:                                "EHGALIKLNHA",
	BOFANBLCKGI:                                "BOFANBLCKGI",
	HGCHMEOGOMA:                                "HGCHMEOGOMA",
	FCPEMJIJNPJ:                                "FCPEMJIJNPJ",
	MENBJHDINMP:                                "MENBJHDINMP",
	FFMDKKCIHMA:                                "FFMDKKCIHMA",
	NBJPOFHICID:                                "NBJPOFHICID",
	NDLOBADOAAO:                                "NDLOBADOAAO",
	LEIBJEKBPBJ:                                "LEIBJEKBPBJ",
	HMJICDOPOLH:                                "HMJICDOPOLH",
	KDCGLMGICMG:                                "KDCGLMGICMG",
	GKMKCEAMKPF:                                "GKMKCEAMKPF",
	OKPOACNHJEK:                                "OKPOACNHJEK",
	AMFBCEOPCDP:                                "AMFBCEOPCDP",
	PABIPMPHMHC:                                "PABIPMPHMHC",
	CCIENKDNOAE:                                "CCIENKDNOAE",
	AGJLOBMFKNC:                                "AGJLOBMFKNC",
	DMGIOHDPGIO:                                "DMGIOHDPGIO",
	BNINEBOOBCA:                                "BNINEBOOBCA",
	DAEDKDPBMIM:                                "DAEDKDPBMIM",
	KAMEILLBJOF:                                "KAMEILLBJOF",
	MIIEPDAIGCH:                                "MIIEPDAIGCH",
	FPEONCPAMOI:                                "FPEONCPAMOI",
	OIHHMIBGPIA:                                "OIHHMIBGPIA",
	GBLKBGALMIB:                                "GBLKBGALMIB",
	PBGEKEEIHCI:                                "PBGEKEEIHCI",
	AEKEJPOHHLO:                                "AEKEJPOHHLO",
	HBPLLJKPEIA:                                "HBPLLJKPEIA",
	NENFNAACLGL:                                "NENFNAACLGL",
	MIHENFCLPHE:                                "MIHENFCLPHE",
	CJGIAPJKIDL:                                "CJGIAPJKIDL",
	AFNHGDHFPGO:                                "AFNHGDHFPGO",
	NACKBGDHDIC:                                "NACKBGDHDIC",
	HGDLDLDKFJD:                                "HGDLDLDKFJD",
	IMNEGOKAGHM:                                "IMNEGOKAGHM",
	GDNGCFMFJIA:                                "GDNGCFMFJIA",
	NIFBFFOBEAN:                                "NIFBFFOBEAN",
	DNKCJELFHOK:                                "DNKCJELFHOK",
	OMMKLGOJJMB:                                "OMMKLGOJJMB",
	EHPCNCMNBOH:                                "EHPCNCMNBOH",
	GMDGFJLKGKF:                                "GMDGFJLKGKF",
	GCAJBFICCCL:                                "GCAJBFICCCL",
	HAPPJBCBMGC:                                "HAPPJBCBMGC",
	IAGEPPHBBPC:                                "IAGEPPHBBPC",
	NODCKLPHMIM:                                "NODCKLPHMIM",
	IKIBCLKMNJJ:                                "IKIBCLKMNJJ",
	MJIMHKNAMCD:                                "MJIMHKNAMCD",
	DNLBFEELOOA:                                "DNLBFEELOOA",
	GNDNNGOMNPB:                                "GNDNNGOMNPB",
	GGKNJFGLKKE:                                "GGKNJFGLKKE",
	OCKFHJAHCCJ:                                "OCKFHJAHCCJ",
	AIPGPJMNOCB:                                "AIPGPJMNOCB",
	FightEnterScRsp:                            "FightEnterScRsp",
	FightHeartBeatScRsp:                        "FightHeartBeatScRsp",
	FightKickOutScNotify:                       "FightKickOutScNotify",
	FightLeaveScNotify:                         "FightLeaveScNotify",
	FightSessionStopScNotify:                   "FightSessionStopScNotify",
	FightGeneralScNotify:                       "FightGeneralScNotify",
	FightGeneralScRsp:                          "FightGeneralScRsp",
	FNDJDJNEAEA:                                "FNDJDJNEAEA",
	OEBPDLGBCBB:                                "OEBPDLGBCBB",
	LIGDMICINAB:                                "LIGDMICINAB",
	IGEIIOKAKLG:                                "IGEIIOKAKLG",
	GBONFNCICAL:                                "GBONFNCICAL",
	BLDPPINBDFC:                                "BLDPPINBDFC",
	FENMJKPFKFO:                                "FENMJKPFKFO",
	FCJHEBMDCHA:                                "FCJHEBMDCHA",
	GGHDPLPGCMN:                                "GGHDPLPGCMN",
	ELNGDGGHLPL:                                "ELNGDGGHLPL",
	JKIAHGLADJF:                                "JKIAHGLADJF",
	CCONBOOCEDK:                                "CCONBOOCEDK",
	CONOIMJDDPB:                                "CONOIMJDDPB",
	BGEMCKFPJFB:                                "BGEMCKFPJFB",
	ICIELFFCAEG:                                "ICIELFFCAEG",
	BCICFMMJJOI:                                "BCICFMMJJOI",
	DFAKFBEFHBI:                                "DFAKFBEFHBI",
	EJJMMHMPGNF:                                "EJJMMHMPGNF",
	BIHJPPJHOHH:                                "BIHJPPJHOHH",
}

func Name(cmdID uint16) string {
	if s, ok := names[cmdID]; ok {
		return s
	}
	return "Cmd(" + itoa(uint32(cmdID)) + ")"
}

func itoa(v uint32) string {
	if v == 0 {
		return "0"
	}
	var buf [10]byte
	i := len(buf)
	for v > 0 {
		i--
		buf[i] = byte('0' + v%10)
		v /= 10
	}
	return string(buf[i:])
}
