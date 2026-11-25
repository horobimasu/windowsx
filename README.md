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
24. `user32.dll!PostMessageW`
25. `user32.dll!SetProcessDPIAware`
26. `user32.dll!GetCursorInfo`
27. `user32.dll!GetIconInfo`
28. `user32.dll!DrawIconEx`
29. `ntdll.dll!RtlSetProcessIsCritical`
30. `ntdll.dll!RtlAdjustPrivilege`
31. `ntdll.dll!NtRaiseHardError`
32. `kernel32.dll!FindFirstStreamW`
33. `kernel32.dll!FindNextStreamW`
34. `kernel32.dll!BeginUpdateResourceW`
35. `kernel32.dll!UpdateResourceW`
36. `kernel32.dll!EndUpdateResourceW`
37. `kernel32.dll!EnumResourceLanguagesW`
38. `kernel32.dll!EnumResourceNamesW`
39. `kernel32.dll!FindResourceW`
40. `kernel32.dll!LoadResource`
41. `kernel32.dll!LockResource`
42. `kernel32.dll!SizeofResource`
43. `kernel32.dll!GlobalAlloc`
44. `kernel32.dll!GlobalFree`
45. `kernel32.dll!GlobalUnlock`
46. `kernel32.dll!GlobalLock`
47. `kernel32.dll!GlobalSize`
48. `kernel32.dll!FreeConsole`
49. `kernel32.dll!AttachConsole`
50. `powrprof.dll!PowerGetActiveScheme`
51. `powrprof.dll!PowerSetActiveScheme`
52. `powrprof.dll!PowerWriteACValueIndex`
53. `powrprof.dll!PowerWriteDCValueIndex`
54. `powrprof.dll!PowerReadACValueIndex`
55. `powrprof.dll!PowerReadDCValueIndex`
56. `gdi32.dll!CreateCompatibleDC`
57. `gdi32.dll!CreateDIBSection`
58. `gdi32.dll!SelectObject`
59. `gdi32.dll!DeleteObject`
60. `gdi32.dll!DeleteDC`
