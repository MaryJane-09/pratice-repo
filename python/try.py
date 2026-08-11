def reverse_string(value):
    text = ""

    for i in range(len(value)-1, -1, -1):
        text += value[i]
    return text
