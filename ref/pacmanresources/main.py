import os

def gen_walls():
    for i in range(1, 25):
        os.system("file2byteslice -package images -input images/{}.png -output images/wall{}.go -var W{}_png".format(i, i, i))
        print("images/wall{}.go".format(i))

def gen_bigdots():
    for i in range(1, 3):
        os.system("file2byteslice -package images -input images/BigDot{}.png -output images/bigdot{}.go -var BigDot{}_png".format(i, i, i))
        print("images/BigDot{}.go".format(i))

def gen_dot():
    os.system("file2byteslice -package images -input images/Dot.png -output images/dot.go -var Dot_png")
    print("images/Dot.go")

def gen_pacman():
    for i in range(1, 9):
        os.system("file2byteslice -package images -input images/pacman{}.png -output images/pacman{}.go -var Pacman{}_png".format(i, i, i))
        print("images/pacman{}.go".format(i))

def gen_blinky():
    for i in range(1, 9):
        os.system("file2byteslice -package images -input images/blinky{}.png -output images/blinky{}.go -var Blinky{}_png".format(i, i, i))
        print("images/blinky{}.go".format(i))

def gen_clyde():
    for i in range(1, 9):
        os.system("file2byteslice -package images -input images/clyde{}.png -output images/clyde{}.go -var Clyde{}_png".format(i, i, i))
        print("images/clyde{}.go".format(i))

def gen_inky():
    for i in range(1, 9):
        os.system("file2byteslice -package images -input images/inky{}.png -output images/inky{}.go -var Inky{}_png".format(i, i, i))
        print("images/inky{}.go".format(i))

def gen_pinky():
    for i in range(1, 9):
        os.system("file2byteslice -package images -input images/pinky{}.png -output images/pinky{}.go -var Pinky{}_png".format(i, i, i))
        print("images/pinky{}.go".format(i))

def gen_vulnerable():
    os.system("file2byteslice -package images -input images/eaten.png -output images/eaten.go -var Eaten_png")
    print("images/eaten.go")

    for i in range(1, 3):
        os.system("file2byteslice -package images -input images/vulnerable{}.png -output images/vulnerable{}.go -var Vulnerable{}_png".format(i, i, i))
        print("images/vulnerable{}.go".format(i))
        os.system("file2byteslice -package images -input images/vulnerableblink{}.png -output images/vulnerableblink{}.go -var Vulnerableblink{}_png".format(i, i, i))
        print("images/vulnerableblink{}.go".format(i))

def gen_fruits():
    for i in range(1, 4):
        os.system("file2byteslice -package images -input images/fruit{}.png -output images/fruit{}.go -var Fruit{}_png".format(i, i, i))
        print("images/fruit{}.go".format(i))

def gen_points():

    for i in points():
        os.system("file2byteslice -package images -input images/{}pts.png -output images/point{}.go -var Point{}_png".format(i, i, i))
        print("images/point{}.go".format(i))

def points():
    c = 100
    l = [c]
    i = 1
    while i <= 4:
        c = c*2
        l.append(c)
        i += 1
    return l

def gen_particles():
    os.system("file2byteslice -package images -input images/PacParticle.png -output images/pacparticle.go -var PacParticle_png")
    print("images/pacparticle.go")
    os.system("file2byteslice -package images -input images/GhostParticle.png -output images/ghostparticle.go -var GhostParticle_png")
    print("images/ghostparticle.go")

def gen_images():
    os.system("file2byteslice -package images -input images/gameover.png -output images/gameover.go -var GameOver_png")
    print("images/gameover.go")
    os.system("file2byteslice -package images -input images/congratulations.png -output images/congratulations.go -var Congrats_png")
    print("images/congratulations.go")

if __name__ == "__main__":
    # gen_walls()
    # gen_bigdots()
    # gen_dot()
    # gen_pacman()
    # gen_blinky()
    # gen_clyde()
    # gen_inky()
    # gen_pinky()
    # gen_vulnerable()
    # gen_fruits()
    # gen_points()
    # gen_particles()
    gen_images()
