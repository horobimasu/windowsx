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
5. `user32.dll!GetMessageW`
6. `user32.dll!PeekMessageW`
7. `user32.dll!PostThreadMessageW`
8. `user32.dll!GetWindowTextW`
9. `user32.dll!GetWindowTextLengthW`
10. `user32.dll!FindWindowW`
11. `user32.dll!SetWindowTextW`
12. `user32.dll!MoveWindow`
13. `user32.dll!SetWindowPos`
14. `user32.dll!GetSystemMetrics`
15. `user32.dll!GetWindowRect`
16. `user32.dll!OpenClipboard`
17. `user32.dll!CloseClipboard`
18. `user32.dll!EmptyClipboard`
19. `user32.dll!SetClipboardData`
20. `user32.dll!GetClipboardData`
21. `user32.dll!RegisterClipboardFormatW`
22. `user32.dll!EnumDisplaySettingsW`
23. `user32.dll!ChangeDisplaySettingsExW`
24. `ntdll.dll!RtlSetProcessIsCritical`
25. `ntdll.dll!RtlAdjustPrivilege`
26. `ntdll.dll!NtRaiseHardError`
27. `kernel32.dll!FindFirstStreamW`
28. `kernel32.dll!FindNextStreamW`
29. `kernel32.dll!BeginUpdateResourceW`
30. `kernel32.dll!UpdateResourceW`
31. `kernel32.dll!EndUpdateResourceW`
32. `kernel32.dll!EnumResourceLanguagesW`
33. `kernel32.dll!EnumResourceNamesW`
34. `kernel32.dll!FindResourceW`
35. `kernel32.dll!LoadResource`
36. `kernel32.dll!LockResource`
37. `kernel32.dll!SizeofResource`
38. `kernel32.dll!GlobalAlloc`
39. `kernel32.dll!GlobalFree`
40. `kernel32.dll!GlobalUnlock`
41. `kernel32.dll!GlobalLock`
42. `kernel32.dll!GlobalSize`
43. `powrprof.dll!PowerGetActiveScheme`
44. `powrprof.dll!PowerSetActiveScheme`
45. `powrprof.dll!PowerWriteACValueIndex`
46. `powrprof.dll!PowerWriteDCValueIndex`
47. `powrprof.dll!PowerReadACValueIndex`
48. `powrprof.dll!PowerReadDCValueIndex`
