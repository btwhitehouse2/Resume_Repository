class Course():
    def __init__(self,BLD,roomNum,maxSize,projector,whiteboard,courseName,minSize,courseNum,sectNum):
        #The following characteristics would be inheritied from RMinBLD class
        #for simplification removed
        self.courseID = int
        self.name = str
        self.BLD = str
        self.room = int
        self.maxSize = int
        self.minSize = int
        self.projector = bool
        self.whiteBoard = bool
        self.Restricted = False
    
    def __str__(self):
        return f"""
        Course: {self.name}
        Loc: {self.BLD} {self.roomNum}
        Num & Sect: {self.courseNum} {self.sectNum}
        Anyone Can Sign Up: {self.Restricted}
    """