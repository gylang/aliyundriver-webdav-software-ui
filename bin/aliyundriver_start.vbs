set ws=WScript.CreateObject("WScript.Shell")
cmd = ""
Set oArgs = WScript.Arguments
    For Each s In oArgs
        cmd = cmd & s & " "
    Next
Set oArgs = Nothing
' MsgBox(Trim(cmd))
ws.Run cmd,0
