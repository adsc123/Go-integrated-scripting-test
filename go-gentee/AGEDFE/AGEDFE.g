run map.int{
	int ioutcome = Arg(`--ioutcome`,0)
	int age = Arg(`--age`,24)
	float totac1 = float(Arg(`--totac1`))
	float acthr = float(Arg(`--acthr`))
	int everot = Arg(`--everot`,0)
	int bacthr
	map.int returnValues

	if ioutcome != 3 {
        if (age >= 0 && age <= 15) {
            bacthr = -9
        } elif totac1 == -9 {
            if acthr == -9 {
                if everot == -8 {
                    bacthr = -8
                } else {
                    bacthr = -9
                }
            } elif acthr == 99 || acthr == -8 {
                bacthr= -8
            } elif acthr > 0 && acthr < 1 {
                bacthr = 1
            } else {
                int workint = int(acthr)
                if (acthr - workint) == 0.5 {
                    if (workint % 2) == 0 {
                        bacthr = workint
                    } else {
                        bacthr = workint + 1
                    }
                } else {
                    bacthr = int(acthr + 0.5)
                }
                if bacthr > 97 {
                    bacthr = 97
                }
            }
        } elif totac1 == 99 || totac1 == -8 {
            bacthr = -8
        } elif totac1 > 0 && totac1 < 1 {
            bacthr = 1
        } else {
            int workint = int(totac1)
            if totac1 - workint == 0.5 {
                if workint % 2 == 0 {
                    bacthr = workint
                } else {
                    bacthr = workint + 1
                }
            } else {
                bacthr = int(totac1 + 0.5)
            }
            if (bacthr > 97) {
                bacthr = 97
            }
        }
    } else {
        bacthr = -9
    }
    returnValues["bacthr"] = bacthr
    return returnValues

}