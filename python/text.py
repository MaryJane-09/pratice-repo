print("=== Text Editing Tool ===")
try:
    input_file = input("Enter the input file: ")
    output_file = input("Enter the output file: ")

    with open(input_file, "r", encoding="utf-8") as file:
        content = file.read()
        print()
        print("The file contains: ")
        print(content)

    with open(output_file, "w", encoding="utf-8") as output:
        output.write(content)
        print("Copy successful")
except FileNotFoundError:
    print("Error: file not found")



