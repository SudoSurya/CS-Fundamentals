# Get user inputs for product prices
a4sheet_price = float(input("Cost of A4sheet:\n"))

# Check if A4 sheet price is negative
if a4sheet_price < 0:
    print("Invalid input")
else:
    pen_price = float(input("Cost of pen:\n"))

    # Check if pen price is negative
    if pen_price < 0:
        print("Invalid input")
    else:
        pencil_price = float(input("Cost of pencil:\n"))

        # Check if pencil price is negative
        if pencil_price < 0:
            print("Invalid input")
        else:
            eraser_price = float(input("Cost of eraser:\n"))

            # Check if eraser price is negative
            if eraser_price < 0:
                print("Invalid input")
            else:
                # Generate and print the price list
                print("\nItems Details\n\nA4sheet:{:.2f}\nPen:{:.2f}\nPencil:{:.2f}\nEraser:{:.2f}".format(a4sheet_price, pen_price, pencil_price, eraser_price))
