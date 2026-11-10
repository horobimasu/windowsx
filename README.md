# windows-extended
go package to add win api functions not in golang.org/x/sys/windows

more functions will be added over time

# install
```bash
go get github.com/horobimasu/windowsx
```

# functions
1. `user32.dll!SetCursorPos`
2. `user32.dll!BlockInput`
3. `user32.dll!SetWinEventHook`
4. `user32.dll!UnhookWinEvent`
5. `user32.dll!GetClassNameW`
6. `user32.dll!GetMessageW`
7. `user32.dll!PeekMessageW`
8. `user32.dll!PostThreadMessageW`
9. `user32.dll!GetWindowTextW`
10. `user32.dll!GetWindowTextLengthW`
11. `user32.dll!FindWindowW`
12. `user32.dll!SetWindowTextW`
13. `user32.dll!MoveWindow`
14. `user32.dll!SetWindowPos`
15. `user32.dll!GetSystemMetrics`
16. `user32.dll!GetWindowRect`
17. `user32.dll!OpenClipboard`
18. `user32.dll!CloseClipboard`
19. `user32.dll!EmptyClipboard`
20. `user32.dll!SetClipboardData`
21. `user32.dll!GetClipboardData`
22. `user32.dll!RegisterClipboardFormatW`
23. `ntdll.dll!RtlSetProcessIsCritical`
24. `ntdll.dll!RtlAdjustPrivilege`
25. `ntdll.dll!NtRaiseHardError`
26. `kernel32.dll!FindFirstStreamW`
27. `kernel32.dll!FindNextStreamW`
28. `kernel32.dll!BeginUpdateResourceW`
29. `kernel32.dll!UpdateResourceW`
30. `kernel32.dll!EndUpdateResourceW`
31. `kernel32.dll!EnumResourceLanguagesW`
32. `kernel32.dll!EnumResourceNamesW`
33. `kernel32.dll!FindResourceW`
34. `kernel32.dll!LoadResource`
35. `kernel32.dll!LockResource`
36. `kernel32.dll!SizeofResource`
37. `kernel32.dll!GlobalAlloc`
38. `kernel32.dll!GlobalFree`
39. `kernel32.dll!GlobalUnlock`
40. `kernel32.dll!GlobalLock`
41. `kernel32.dll!GlobalSize`
