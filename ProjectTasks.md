# Project Tasks

Maintained by the **project-manager** agent (see `.claude/agents/project-manager.md`). Every task performed in this repo — whether requested through Claude Code or discovered as a manual check-in — gets a row here.

Do not hand-edit the Task ID counter below; the agent reads the last-used ID and increments it.

<!-- next-task-id: TASK-0003 -->

| Task ID | Task Summary | Status | Start Date | End Date | Remarks | Created By |
|---------|--------------|--------|------------|----------|---------|------------|
| TASK-0001 | Bootstrap project-manager agent, ProjectTasks.md ledger, and manual-commit reconciliation hook | Finished | 2026-08-07 | 2026-08-07 | Rolled out from the employee.unirms.com pilot. Committed directly (not via project-manager) since the agent didn't exist yet to commit its own bootstrap. | Ashutosh Dhiman |
| TASK-0002 | Add `Timestamp` field to `USBDeviceInfo` and new `USBDeviceInfos []USBDeviceInfo` array field to `ActivityReport` in `hrms/activity_monitor.go`, so `activity.monitor.unirms` has somewhere to persist USB connect/disconnect events it already receives but currently discards | Finished | 2026-09-15 | 2026-09-15 | Purely additive (new bson tags `timestamp`, `usb_device_infos`); no existing writer of `usb_device_info` sub-document per developer's DB inspection, so no migration concern. Positional-struct-literal check: repo-wide grep for `USBDeviceInfo{`/`ActivityReport{` found zero `hrms.USBDeviceInfo{`/`hrms.ActivityReport{` literals — the only 3 `USBDeviceInfo{` hits (genproto/monitor/activity_log.pb.go:573, grpcclient/monitor_requester.go:163,304) construct the unrelated protobuf type `monitor.USBDeviceInfo`, all keyed — so field-order shift is safe. gofmt/build/vet all clean (verified). Reviewed inline against correctness/security/readability/dead-code standard — no Task/subagent tool available in this session to delegate to go-reviewer; no CRITICAL/HIGH findings. Not committed — git operations stay with developer. | Ashutosh Dhiman |
