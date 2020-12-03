#!/usr/bin/env python3

"""
    flip - Jonah G. Rongstad (2020)
"""

from curses import *
from sys import argv, stderr

s = initscr()
noecho()
s.keypad(1)
curs_set(0)

def main():
    if len(argv) == 1:
        endwin()

        print("You must supply arguments!", file=stderr)

        exit()

    height, width = s.getmaxyx()

    cards = get_files()

    if cards == IOError:
        endwin()

        print("Invalid file name(s)!", file=stderr)

        exit()

    card = 0
    display_bar = True;

    disp_file(cards[card])
    disp_bar(card + 1, len(cards))

    s.refresh()

    while True:
        nh, nw = s.getmaxyx()

        if height != nh or width != nw:
            height = nh
            width = nw

            s.clear()
            disp_file(cards[card])
            disp_bar(card + 1, len(cards))

            s.refresh()

        ch = s.getch()

        if ch == ord("q"):
            endwin()
            exit()
        elif ch == ord("l") or ch == KEY_RIGHT:
            card = roll_over(card, len(cards) - 1, 0, 1)
            draw_screen(cards[card])
        elif ch == ord("h") or ch == KEY_LEFT:
            card = roll_over(card, 0, len(cards) - 1, -1)
            draw_screen(cards[card])
        elif ch == ord("b"):
            display_bar = not display_bar

            if not display_bar:
                s.addstr(height-1, 0, ' ' * len(str(card+1) + "/" + str(len(cards))))

        if display_bar:
            disp_bar(card + 1, len(cards))

def draw_screen(file):
    s.clear()
    disp_file(file)

    s.refresh()

def roll_over(c, edge, roll, add):
    if c == edge:
        return roll
    else:
        return add

def disp_file(lines):
    height, _ = s.getmaxyx()

    for i, l in enumerate(lines):
        if i + 1 < height:
            s.addstr(l)
            s.move(i + 1, 0)
        else:
            break

def disp_bar(index, total):
    height, _ = s.getmaxyx()

    s.addstr(height - 1, 0, str(index) + "/" + str(total), A_REVERSE)

def get_files():
    cards = []

    for fname in argv[1:]:
        try:
            with open(fname, "r") as f:
                cards.append(f.readlines())
        except:
            return IOError

    return cards

if __name__ == "__main__":
    main()
