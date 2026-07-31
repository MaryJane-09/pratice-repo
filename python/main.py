# "Hello, [Name]! Your [Drink] is ready."

name = input("Input your name: ")
cleaned_name = name.capitalize()
if cleaned_name.isalpha():
    drink = input("Input your drink: ")
    cleaned_drink = drink.strip().capitalize()
    print(f"Hello, {cleaned_name}! Your {cleaned_drink} is ready.")
else:
    print("Name must contain letters only")
