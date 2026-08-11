def safe_calculator(a, operator, b):
    operators = {"+", "-", "*", "/", "%", "**"}

    int1 = float(a)
    int2 = float(b)

    if operator not in operators:
        return "Invalid operator"

    if (operator == "/" or operator == "%") and int2 == 0:
            return "Cannot divide by zero"
    
    if operator == "+":
         return round(int1, 2) + round(int2, 2)
    if operator == "-":
             return round(int1, 2) - round(int2, 2)
    if operator == "*":
             return round(int1, 2) * round(int2, 2)
    if operator == "/":
             return round(int1, 2) / round(int2, 2)
    if operator == "%":
             return round(int1, 2) % round(int2, 2)
    if operator == "**":
             return round(int1, 2) ** round(int2, 2)
    
print(safe_calculator(5, "+", 3))