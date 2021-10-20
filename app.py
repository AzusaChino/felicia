import cv2


def main():
    file_path = "./peek.jpg"
    img = cv2.imread(file_path)
    # turn the img to gray
    gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)

    # draw rectangles
    x = y = 10  # coordinate axis
    w = 100  # rect area
    color = (0, 0, 255)
    cv2.rectangle(img, (x, y), (x + w, y + w), color, 1)
    cv2.imshow("Image", img)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
