Sample from : http://software-solutions-online.com/2014/03/13/excel-vba-save-file-dialog-getsaveasfilename/


Sub Example1()
Dim objFSO As Object
Dim objFolder As Object
Dim objFile As Object
Dim i As Integer 

'Create an instance of the FileSystemObject 
Set objFSO = CreateObject("Scripting.FileSystemObject")
'Get the folder object 
Set objFolder = objFSO.GetFolder("D:\Stuff\Freelances\Website\Blog\Arrays\Pics")
i = 1
'loops through each file in the directory and prints their names and path 
For Each objFile In objFolder.Files
    'print file name 
    Cells(i + 1, 1) = objFile.Name 
    'print file path 
    Cells(i + 1, 2) = objFile.Path 
    i = i + 1 
Next objFile
End Sub 


'Using the code below, the names of the folders and their associated paths are listed on column A and B:
Sub Example2()
Dim objFSO As Object
Dim objFolder As Object
Dim objSubFolder As Object
Dim i As Integer 

'Create an instance of the FileSystemObject 
Set objFSO = CreateObject("Scripting.FileSystemObject")
'Get the folder object 
Set objFolder = objFSO.GetFolder("D:\Stuff\Freelances\Website\Blog")
i = 1
'loops through each file in the directory and prints their names and path 
For Each objSubFolder In objFolder.subfolders
    'print folder name 
    Cells(i + 1, 1) = objSubFolder.Name 
    'print folder path 
    Cells(i + 1, 2) = objSubFolder.Path 
    i = i + 1 
Next objSubFolder
End Sub 





