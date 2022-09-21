import os
 
# Get the list of all files and directories
path = "./"
dir_list = os.listdir(path)
 
print("Files and directories in '", path, "' :")

# put list of files dir_list in a file named "list.txt"
with open("list.txt", "w") as f:
    for file in dir_list:
        f.write(file + "\n")
        # print(file)

# read the file "list.txt" and print the content
# with open("list.txt", "r") as f:
#     for line in f:
#         print(line, end="")
 