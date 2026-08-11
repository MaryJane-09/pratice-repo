
def bubble_sort(arr):
    n = len(arr)
    for i in range(n):
        for j in range(0, n - i - 1):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]

prices = [4.50, 3.50, 5.00, 2.00]
bubble_sort(prices)
print(prices)

def selection_sort(arr):
    n = len(arr)
    for i in range(n):
        min_index = i  # Assume the current item is the smallest
        for j in range(i + 1, n):
            if arr[j] < arr[min_index]:
                min_index = j  # Update the index of the smallest item found
                
        # Swap the smallest item with the first unsorted item
        arr[i], arr[min_index] = arr[min_index], arr[i]

price = [4.50, 3.50, 5.00, 2.00]
selection_sort(price)
print(price)

def insertion_sort(arr):
    for i in range(1, len(arr)):
        key = arr[i]  # Pick up the current cup
        j = i - 1
        
        # Slide the cup left as long as the neighboring cup is more expensive
        while j >= 0 and arr[j] > key:
            arr[j + 1] = arr[j]  # Shift the more expensive cup to the right
            j -= 1
            
        arr[j + 1] = key  # Drop the cup into its correct spot

prices = [4.50, 3.50, 5.00, 2.00]
insertion_sort(prices)
print(prices)