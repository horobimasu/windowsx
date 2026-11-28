# windows-extended
go package to add win api functions not in golang.org/x/sys/windows

more functions will be added over time

# install
```bash
go get github.com/horobimasu/windowsx
```

# functions
1. `user32.dll!SetCursorPos`
2. `user32.dll!GetCursorPos`
3. `user32.dll!BlockInput`
4. `user32.dll!SetWinEventHook`
5. `user32.dll!UnhookWinEvent`
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
23. `user32.dll!EnumDisplaySettingsW`
24. `user32.dll!ChangeDisplaySettingsExW`
25. `user32.dll!PostMessageW`
26. `user32.dll!SetProcessDPIAware`
27. `user32.dll!GetCursorInfo`
28. `user32.dll!GetIconInfo`
29. `user32.dll!DrawIconEx`
30. `user32.dll!GetWindowDisplayAffinity`
31. `user32.dll!SetWindowDisplayAffinity`
32. `user32.dll!SendInput`
33. `user32.dll!RegisterClassExW`
34. `user32.dll!CreateWindowExW`
35. `user32.dll!DefWindowProcW`
36. `user32.dll!ShowWindow`
37. `user32.dll!UpdateWindow`
38. `user32.dll!DestroyWindow`
39. `ntdll.dll!RtlSetProcessIsCritical`
40. `ntdll.dll!RtlAdjustPrivilege`
41. `ntdll.dll!NtRaiseHardError`
42. `kernel32.dll!FindFirstStreamW`
43. `kernel32.dll!FindNextStreamW`
44. `kernel32.dll!BeginUpdateResourceW`
45. `kernel32.dll!UpdateResourceW`
46. `kernel32.dll!EndUpdateResourceW`
47. `kernel32.dll!EnumResourceLanguagesW`
48. `kernel32.dll!EnumResourceNamesW`
49. `kernel32.dll!FindResourceW`
50. `kernel32.dll!LoadResource`
51. `kernel32.dll!LockResource`
52. `kernel32.dll!SizeofResource`
53. `kernel32.dll!GlobalAlloc`
54. `kernel32.dll!GlobalFree`
55. `kernel32.dll!GlobalUnlock`
56. `kernel32.dll!GlobalLock`
57. `kernel32.dll!GlobalSize`
58. `kernel32.dll!FreeConsole`
59. `kernel32.dll!AttachConsole`
60. `powrprof.dll!PowerGetActiveScheme`
61. `powrprof.dll!PowerSetActiveScheme`
62. `powrprof.dll!PowerWriteACValueIndex`
63. `powrprof.dll!PowerWriteDCValueIndex`
64. `powrprof.dll!PowerReadACValueIndex`
65. `powrprof.dll!PowerReadDCValueIndex`
66. `gdi32.dll!CreateCompatibleDC`
67. `gdi32.dll!CreateDIBSection`
68. `gdi32.dll!SelectObject`
69. `gdi32.dll!DeleteObject`
70. `gdi32.dll!DeleteDC`
71. `gdi32.dll!CreateSolidBrush`
