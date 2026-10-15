package windowsx

const (
	SE_CREATE_TOKEN_PRIVILEGE                      uint32 = 2
	SE_ASSIGNPRIMARYTOKEN_PRIVILEGE                uint32 = 3
	SE_LOCK_MEMORY_PRIVILEGE                       uint32 = 4
	SE_INCREASE_QUOTA_PRIVILEGE                    uint32 = 5
	SE_MACHINE_ACCOUNT_PRIVILEGE                   uint32 = 6
	SE_TCB_PRIVILEGE                               uint32 = 7
	SE_SECURITY_PRIVILEGE                          uint32 = 8
	SE_TAKE_OWNERSHIP_PRIVILEGE                    uint32 = 9
	SE_LOAD_DRIVER_PRIVILEGE                       uint32 = 10
	SE_SYSTEM_PROFILE_PRIVILEGE                    uint32 = 11
	SE_SYSTEMTIME_PRIVILEGE                        uint32 = 12
	SE_PROF_SINGLE_PROCESS_PRIVILEGE               uint32 = 13
	SE_INC_BASE_PRIORITY_PRIVILEGE                 uint32 = 14
	SE_CREATE_PAGEFILE_PRIVILEGE                   uint32 = 15
	SE_CREATE_PERMANENT_PRIVILEGE                  uint32 = 16
	SE_BACKUP_PRIVILEGE                            uint32 = 17
	SE_RESTORE_PRIVILEGE                           uint32 = 18
	SE_SHUTDOWN_PRIVILEGE                          uint32 = 19
	SE_DEBUG_PRIVILEGE                             uint32 = 20
	SE_AUDIT_PRIVILEGE                             uint32 = 21
	SE_SYSTEM_ENVIRONMENT_PRIVILEGE                uint32 = 22
	SE_CHANGE_NOTIFY_PRIVILEGE                     uint32 = 23
	SE_REMOTE_SHUTDOWN_PRIVILEGE                   uint32 = 24
	SE_UNDOCK_PRIVILEGE                            uint32 = 25
	SE_SYNC_AGENT_PRIVILEGE                        uint32 = 26
	SE_ENABLE_DELEGATION_PRIVILEGE                 uint32 = 27
	SE_MANAGE_VOLUME_PRIVILEGE                     uint32 = 28
	SE_IMPERSONATE_PRIVILEGE                       uint32 = 29
	SE_CREATE_GLOBAL_PRIVILEGE                     uint32 = 30
	SE_TRUSTED_CREDMAN_ACCESS_PRIVILEGE            uint32 = 31
	SE_RELABEL_PRIVILEGE                           uint32 = 32
	SE_INC_WORKING_SET_PRIVILEGE                   uint32 = 33
	SE_TIME_ZONE_PRIVILEGE                         uint32 = 34
	SE_CREATE_SYMBOLIC_LINK_PRIVILEGE              uint32 = 35
	SE_DELEGATE_SESSION_USER_IMPERSONATE_PRIVILEGE uint32 = 36
)

const (
	EVENT_AIA_START                               uint32 = 0xA000
	EVENT_AIA_END                                 uint32 = 0xAFFF
	EVENT_MIN                                     uint32 = 0x00000001
	EVENT_MAX                                     uint32 = 0x7FFFFFFF
	EVENT_OBJECT_ACCELERATORCHANGE                uint32 = 0x8012
	EVENT_OBJECT_CLOAKED                          uint32 = 0x8017
	EVENT_OBJECT_CONTENTSCROLLED                  uint32 = 0x8015
	EVENT_OBJECT_CREATE                           uint32 = 0x8000
	EVENT_OBJECT_DEFACTIONCHANGE                  uint32 = 0x8011
	EVENT_OBJECT_DESCRIPTIONCHANGE                uint32 = 0x800D
	EVENT_OBJECT_DESTROY                          uint32 = 0x8001
	EVENT_OBJECT_DRAGSTART                        uint32 = 0x8021
	EVENT_OBJECT_DRAGCANCEL                       uint32 = 0x8022
	EVENT_OBJECT_DRAGCOMPLETE                     uint32 = 0x8023
	EVENT_OBJECT_DRAGENTER                        uint32 = 0x8024
	EVENT_OBJECT_DRAGLEAVE                        uint32 = 0x8025
	EVENT_OBJECT_DRAGDROPPED                      uint32 = 0x8026
	EVENT_OBJECT_END                              uint32 = 0x80FF
	EVENT_OBJECT_FOCUS                            uint32 = 0x8005
	EVENT_OBJECT_HELPCHANGE                       uint32 = 0x8010
	EVENT_OBJECT_HIDE                             uint32 = 0x8003
	EVENT_OBJECT_HOSTEDOBJECTSINVALIDATED         uint32 = 0x8020
	EVENT_OBJECT_IME_HIDE                         uint32 = 0x8028
	EVENT_OBJECT_IME_SHOW                         uint32 = 0x8027
	EVENT_OBJECT_IME_CHANGE                       uint32 = 0x8029
	EVENT_OBJECT_INVOKED                          uint32 = 0x8013
	EVENT_OBJECT_LIVEREGIONCHANGED                uint32 = 0x8019
	EVENT_OBJECT_LOCATIONCHANGE                   uint32 = 0x800B
	EVENT_OBJECT_NAMECHANGE                       uint32 = 0x800C
	EVENT_OBJECT_PARENTCHANGE                     uint32 = 0x800F
	EVENT_OBJECT_REORDER                          uint32 = 0x8004
	EVENT_OBJECT_SELECTION                        uint32 = 0x8006
	EVENT_OBJECT_SELECTIONADD                     uint32 = 0x8007
	EVENT_OBJECT_SELECTIONREMOVE                  uint32 = 0x8008
	EVENT_OBJECT_SELECTIONWITHIN                  uint32 = 0x8009
	EVENT_OBJECT_SHOW                             uint32 = 0x8002
	EVENT_OBJECT_STATECHANGE                      uint32 = 0x800A
	EVENT_OBJECT_TEXTEDIT_CONVERSIONTARGETCHANGED uint32 = 0x8030
	EVENT_OBJECT_TEXTSELECTIONCHANGED             uint32 = 0x8014
	EVENT_OBJECT_UNCLOAKED                        uint32 = 0x8018
	EVENT_OBJECT_VALUECHANGE                      uint32 = 0x800E
	EVENT_OEM_DEFINED_START                       uint32 = 0x0101
	EVENT_OEM_DEFINED_END                         uint32 = 0x01FF
	EVENT_SYSTEM_ALERT                            uint32 = 0x0002
	EVENT_SYSTEM_ARRANGMENTPREVIEW                uint32 = 0x8016
	EVENT_SYSTEM_CAPTUREEND                       uint32 = 0x0009
	EVENT_SYSTEM_CAPTURESTART                     uint32 = 0x0008
	EVENT_SYSTEM_CONTEXTHELPEND                   uint32 = 0x000D
	EVENT_SYSTEM_CONTEXTHELPSTART                 uint32 = 0x000C
	EVENT_SYSTEM_DESKTOPSWITCH                    uint32 = 0x0020
	EVENT_SYSTEM_DIALOGEND                        uint32 = 0x0011
	EVENT_SYSTEM_DIALOGSTART                      uint32 = 0x0010
	EVENT_SYSTEM_DRAGDROPEND                      uint32 = 0x000F
	EVENT_SYSTEM_DRAGDROPSTART                    uint32 = 0x000E
	EVENT_SYSTEM_END                              uint32 = 0x00FF
	EVENT_SYSTEM_FOREGROUND                       uint32 = 0x0003
	EVENT_SYSTEM_MENUPOPUPEND                     uint32 = 0x0007
	EVENT_SYSTEM_MENUPOPUPSTART                   uint32 = 0x0006
	EVENT_SYSTEM_MENUEND                          uint32 = 0x0005
	EVENT_SYSTEM_MENUSTART                        uint32 = 0x0004
	EVENT_SYSTEM_MINIMIZEEND                      uint32 = 0x0017
	EVENT_SYSTEM_MINIMIZESTART                    uint32 = 0x0016
	EVENT_SYSTEM_MOVESIZEEND                      uint32 = 0x000B
	EVENT_SYSTEM_MOVESIZESTART                    uint32 = 0x000A
	EVENT_SYSTEM_SCROLLINGEND                     uint32 = 0x0013
	EVENT_SYSTEM_SCROLLINGSTART                   uint32 = 0x0012
	EVENT_SYSTEM_SOUND                            uint32 = 0x0001
	EVENT_SYSTEM_SWITCHEND                        uint32 = 0x0015
	EVENT_SYSTEM_SWITCHSTART                      uint32 = 0x0015
	EVENT_UIA_EVENTID_START                       uint32 = 0x4E00
	EVENT_UIA_EVENTID_END                         uint32 = 0x4EFF
	EVENT_UIA_PROPID_START                        uint32 = 0x7500
	EVENT_UIA_PROPID_END                          uint32 = 0x75FF
)

const (
	WINEVENT_OUTOFCONTEXT   uint32 = 0x0000
	WINEVENT_SKIPOWNTHREAD  uint32 = 0x0001
	WINEVENT_SKIPOWNPROCESS uint32 = 0x0002
	WINEVENT_INCONTEXT      uint32 = 0x0004
)

const (
	PM_NOREMOVE uint32 = 0x0000
	PM_REMOVE   uint32 = 0x0001
	PM_NOYIELD  uint32 = 0x0002
)
