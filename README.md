# windows-extended
go package to add win api functions not in golang.org/x/sys/windows

more functions will be added over time

# install
```bash
go get github.com/horobimasu/windowsx
```

# functions
1. `user32!SetCursorPos`
2. `user32!BlockInput`
3. `user32!SetEventHook`
4. `user32!UnhookEvent`
5. `user32!GetClassName`
6. `user32!GetMessage`
7. `user32!PeekMessage`
8. `user32!PostThreadMessage`
9. `ntdll!RtlSetProcessIsCritical`
10. `ntdll!NtRaiseHardError`
11. `ntdll!RtlAdjustPrivilege`
