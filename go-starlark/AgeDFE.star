def AgeDFE():
    if ioutcome != 3:
        if (age >= 0 and  age <= 15):
            bacthr = -9
        elif totac1 == -9:
            if acthr == -9:
                if everot == -8:
                    bacthr = -8
                else:
                    bacthr = -9
            elif acthr == 99 or acthr == -8:
                bacthr = -8
            elif acthr > 0 and acthr < 1:
                bacthr = 1
            else:
                workint = int(acthr)
                if acthr - workint == 0.5:
                    if workint % 2 == 0:
                        bacthr = workint
                    else:
                        bacthr = workint + 1
                else:
                    bacthr = int(acthr + 0.5)
                if bacthr > 97:
                    bacthr = 97
        elif totac1 == 99 or totac1 == -8:
            bacthr = -8
        elif totac1 > 0 and totac1 < 1:
            bacthr = 1
        else:
            workint = int(totac1)
            if totac1 - workint == 0.5:
                if workint % 2 == 0:
                    bacthr = workint
                else:
                    bacthr = workint + 1
            else:
                bacthr = int(totac1 + 0.5)
            if bacthr > 97:
                bacthr = 97
    else:
        bacthr = -9
    return {"bacthr": bacthr}