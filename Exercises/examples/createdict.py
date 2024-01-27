# Function to create and display the dictionary
def create_dictionary():
    word_dictionary = {}

    while True:
        word = input("Enter the word: ")
        num_of_meanings = int(input("Enter the no of meanings: "))

        # Check if the no. of meanings is zero or negative
        if num_of_meanings <= 0:
            print("Invalid Input")
            if word_dictionary:
                print("\nHere's the dictionary you've created:")
                for key, values in word_dictionary.items():
                    print(f"{key} : {values}\n")
            break

        meanings = []
        print("Enter the meanings :\n")
        for _ in range(num_of_meanings):
            meaning = input()
            meanings.append(meaning)

        word_dictionary[word] = meanings

        choice = input("Do you want to add one more element to the dictionary? If yes, press 1, else press 0: ")

        if choice == '0':
            print("\nHere's the dictionary you've created:")
            for key, values in word_dictionary.items():
                print(f"{key} : {values}\n")
            break
        elif choice != '1':
            print("Invalid Input")
            print("\nHere's the dictionary you've created:")
            for key, values in word_dictionary.items():
                print(f"{key} : {values}\n")
            break

# Call the function to create and display the dictionary
create_dictionary()
