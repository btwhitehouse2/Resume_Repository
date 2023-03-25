from ctypes import Union
import mysql.connector

registrationIsOpen = False

db = mysql.connector.connect(
    host="localhost",
    user="root",
    passwd="scoutwhitehouse",
    database="regie"
)

mycursor = db.cursor()

class UserInterface:
    def recieveCourseRequest(userid,courseid,firstName,lastName)->str:
        ut = UserTableInterface
        ct = CourseTableInterface
        lt = LogTableInterface
        if registrationIsOpen == True:
            if True == ut.isStudent(userid):
                if True == ct.isRestricted(courseid):
                    #has not applied
                    if None == lt.hasNotRequest(userid,courseid):
                        if True == ut.hasRoomInSchedule(userid):
                            if False == ct.isFull(courseid): 
                                messageToFac = UserInterface.sendRestrictRequestToFac(userid,courseid,firstName,lastName)
                                lt.submitRestCourseReq(userid,courseid,"waiting")
                                messageToStud = f"""Hi, {firstName} we have reiceved your request for {courseid} because this class is restricted you will have to wait to be approved and status will be reflected as waiting"""
                                return messageToFac,messageToStud
                            else:
                                return f"""Sorry, {firstName} the following course id: {courseid} is full please select another course to take."""
                        #student has registered for maximum number of classes
                        elif False == ut.hasRoomInSchedule(userid): 
                            return f"""Sorry, {firstName} we have you taking the maximum number of allowed classes for this quarter."""
                            
                    #has applied and still waiting for faculty response
                    elif True == lt.hasRequestWait(courseid,userid): 
                        return  f"""Hi, {firstName} our records show that you have already applied for {courseid} and we are currerntly waiting to here if you have been approved."""
                    #has applied and been rejected
                    elif True == lt.hasRequestDenied(userid,courseid): 
                        return f"""Hi, {firstName} we have reiceved your request for {courseid} we are sorry, but you have been denied from this course for this quarter."""
                    #has already been approved
                    elif True == lt.hasRequestApproved(userid,courseid): 
                        return f"""Hi, {firstName} we have reiceved your request for {courseid} our records show that you have already requested this course and been approved, no need to apply again."""
                elif False == ct.isRestricted(courseid): 
                    
                    #regular class with no restrictions
                    if True == ut.hasRoomInSchedule(userid):
                        if False == ct.isFull(courseid): 
                            lt.submitRegCourseReq(userid,courseid,"registered")
                            return f"""Hi, {firstName} we have reiceved your request for {courseid} and you have been successfully registered"""
                        else:
                            return f"""Sorry, {firstName} the following course id: {courseid} is full please select another course to take."""
                    elif False == ut.hasRoomInSchedule(userid): #student has registered for maximum number of classes
                        return f"""Sorry, {firstName} we have you taking the maximum number of allowed classes for this quarter."""

                
                #incorrect requests
                else: #CourseID is not valid will result in None not True or False
                    return f"""Sorry, {firstName} the following course id: {courseid} is not a valid id."""
            elif False == ut.isStudent(userid): #user that is not a student ie. admin or faculty
                return f"""Sorry, {firstName} we currently do not have you as a student and are unable to register for courses at this time."""
            else:#if submit non valid user ID will result in None not true or false
                return f"""Sorry, {firstName} the following user id: {courseid} is not a valid id."""
        else: #registration is not open
            return f"""Sorry, {firstName} course registration is not currently open"""; 

    def OpenRegistration(userid,firstname,lastname)->str:
        global registrationIsOpen
        if True == UserTableInterface.isAdmin(userid):
            if registrationIsOpen == False:
                registrationIsOpen = True
                return f"""Thank you {firstname} {lastname} for your request, registration is now open"""
            else:
                return f"""Thank you {userid} for your request, but registration is already open"""
        elif False == UserTableInterface.isAdmin(userid):
            return f"""Thank you {userid} for your request, but you do not have admin rights"""
        else:
            return f"""Thank you {userid} for your request, but your user id is not valid"""

    def CloseRegistration(self,userid)->str:
        if True == UserTableInterface.isAdmin(userid):
            if self.isRegOpen == True:
                self.isRegOpen = False
                return f"""Thank you {userid} for your request, registration is now closed"""
            else:
                return f"""Thank you {userid} for your request, but registration was already closed"""
        elif False == UserTableInterface.isAdmin(userid):
                return f"""Thank you {userid} for your request, but you do not have admin rights"""
        else:
            return f"""Thank you {userid} for your request, but your user id is not valid"""

    def sendRestrictRequestToFac(userid,courseid,studfirst,studlast)->str:
        studentFull = studfirst + " " + studlast
        coursename, facultyid = CourseTableInterface.returnCourseInfo(courseid) 
        email, facultyName = UserTableInterface.returnFacultyInfo(facultyid)
        message = f"""Hello {facultyName} the following student: {studentFull} is requesting to take your {coursename} with courseID: {courseid}. Please respond back with accept or decline'\n' """
        return message 

    def sendCurrentCourseRegistered(userid,firstname)->tuple[str,str]:
        listCourseIDWait = LogTableInterface.studReqWaitCourse(userid)
        if len(listCourseIDWait) == 0:
            waitMessage = f"""Hi {firstname} you are not currently waiting for any courses"""
        else:
            waitCourseNames = CourseTableInterface.returnCourseNames(listCourseIDWait)
            waitMessage =  f"""Hi {firstname} you are waiting on approval for the following: {waitCourseNames}"""
        listCourseIDReg = LogTableInterface.studReqRegCourse(userid)
        if len(listCourseIDReg) == 0:
            regMessage = f"""Hi {firstname} you are not currently fully registered for any courses"""
        else:
            regCourseNames = CourseTableInterface.returnCourseNames(listCourseIDReg)
            regMessage =  f"""Hi {firstname} you are currently registered for the following: {regCourseNames}"""    
        return waitMessage, regMessage

    def RecieveFacultyResponse(courseID,studentID,facultyID,acceptStat,facFirst)->str:
        #check to see if this faculty member is correct for the course
        if True == CourseTableInterface.isFacForCourse(courseID,facultyID):
            if True == LogTableInterface.hasRequestWait(studentID,courseID):
                logID = LogTableInterface.lookUpLogID(studentID,courseID)
                if acceptStat == True:
                    LogTableInterface.updateLogStatus(logID,"registered")
                    return f""" Hi {facFirst} you have successfuly approved student with id: {studentID} for your course """
                elif acceptStat == False:
                    LogTableInterface.updateLogStatus(logID,"denied")
                    return f""" Hi {facFirst} you have successfuly denied student with id: {studentID} for your course """
            else:
                return f"""Hi {facFirst} we do not have student with id number: {studentID} currently waiting for course with id: {courseID} """
        else:
            return f"""Hi {facFirst} we do not have you as the professor for course id: {courseID} """

    def StatusFacCourse(facID,facFirstName)->list[str]:
        listMessage = []
        IntroMess = f"""Hi {facFirstName} the status for the following course is as follows: """
        listMessage.append(IntroMess)
        listIds, listNames = CourseTableInterface.returnCourseIDsNames(facID)
        if len(listNames) == 0:
            return f"""Hi {facFirstName} you are not currently instructing any courses"""
        else:
            for i in range(len(listNames)):
                id = listIds[i]
                coursename = listNames[i]
                mSize,cSize = CourseTableInterface.courseStats(id)
                message = f"""{coursename} with id {id} is currently at {cSize} of {mSize}"""
                listMessage.append(message)
            return listMessage
    
    def NotifyAdminOfProfRequest(coursename,professorID,restricted,facFirst,faclast,AdminId)->str:
        AdminName = UserTableInterface.getAdminName(AdminId)
        if restricted == True:
            return f"""Hi {AdminName}, Professor {facFirst} {faclast} with ID {professorID} is requesting a new restricted course called {coursename} """
        else:
            return f"""Hi {AdminName}, Professor {facFirst} {faclast} with ID {professorID} is requesting a new standard course called {coursename} """

    def SubmitNewCourse(adminID,AdminFirst,AdminLast,courseID,courseName,professorID,bld,roomNum,maxSize,minSize,projector,whiteboard,restricted)->str:
        if UserTableInterface.isAdmin(adminID) == True:
            CourseTableInterface.createCourse(courseID,courseName,professorID,bld,roomNum,maxSize,minSize,projector,whiteboard,restricted)
            return f"""New Course: {courseName} with {courseID} was successfully added to the course registration"""
        else:    
            return f"""Hi {AdminFirst} {AdminLast} we currently do not have your account listed as an administrator"""

class UserTableInterface:
    def isStudent(id)->bool:
        mycursor.execute("SELECT role FROM regie.user WHERE userid =%s",[id])
        for x in mycursor:
            if x[0] == "student":
                return True
            else:
                return False
    
    def isAdmin(id)->bool:
        mycursor.execute("SELECT role FROM regie.user WHERE userid =%s",[id])
        for x in mycursor:
            if x[0] == "admin":
                return True
            else:
                return False

    def isFullTime(id)->bool:
        mycursor.execute("SELECT status FROM regie.user WHERE userid =%s",[id])      
        for x in mycursor:
            if x[0] == "full":
                return True
            else:
                return False

    def returnFacultyInfo(professorid)->tuple[str,str]:
        mycursor.execute("SELECT userfirst, userlast, email from regie.user where userid =%s",[professorid])
        for x in mycursor:
            first = x[0]
            last = x[1]
            email = x[2]
            facName = first+" "+last
            return email, facName
    
    def getAdminName(adminid)->str:
        mycursor.execute("SELECT userfirst, userlast from regie.user where userid =%s",[adminid])
        for x in mycursor:
            first = x[0]
            last = x[1]
            adName = first+" "+last
            return adName
    
    def hasRoomInSchedule(userid)->bool:
        numCourse = LogTableInterface.studNumRegCourse(userid)
        status = UserTableInterface.isFullTime(userid)
        if status == True:
            if numCourse < 3:
                return True
            else:
                return False    
        elif status == False:
            if numCourse < 1:
                return True
            else:
                return False
        
    def returnUserInfo(userid)->tuple[str,str,str,str,str,str]:
        mycursor.execute("SELECT userid, userfirst, userlast, email, username, role, status from regie.user where userid =%s",[userid])
        for x in mycursor:
            id = x[0]
            first = x[1]
            last = x[2]
            email = x[3]
            username = x[4]
            role = x[5]
            status = x[6]
            fullName = first+" "+last
            return id, fullName, email, username, role, status
        

class CourseTableInterface:
    def isRestricted(id)->bool:
        mycursor.execute("SELECT restricted from regie.course where courseid =%s",[id])
        for x in mycursor:
            return x[0] 

    def isFull(courseid) ->bool:
        mSize = CourseTableInterface.maxSize(courseid)
        currSize = LogTableInterface.numRegAndWait(courseid)
        if currSize >= mSize:
            return True
        else:
            return False

    def courseStats(courseid)->tuple[int,int]:
        mSize = CourseTableInterface.maxSize(courseid)
        currSize = LogTableInterface.numRegAndWait(courseid)
        return mSize, currSize

    
    def isFacForCourse(courseid,professorid)->bool:
        mycursor.execute("SELECT professorid FROM regie.course WHERE courseid =%s",[courseid])      
        for x in mycursor:
            if x[0] == professorid:
                return True
            else:
                return False

    def maxSize(courseid) -> int:
        mycursor.execute("SELECT maxsize FROM regie.course where courseid =%s",[courseid])
        for x in mycursor:
            return(x[0])

    def returnCourseInfo(courseid)->tuple[str,int]:
        mycursor.execute("SELECT coursename, professorid from regie.course where courseid =%s",[courseid])
        for x in mycursor:
            coursename = x[0]
            professorid = x[1]
            return coursename, professorid
    
    def returnCourseNames(courseids)->list[str]:
        listCourseNames = []
        for i in range(len(courseids)): 
            mycursor.execute("SELECT coursename from regie.course where courseid =%s",[courseids[i]])
            for x in mycursor:
                coursename = x[0]
                listCourseNames.append(coursename)
        return listCourseNames
    
    def returnCourseIDsNames(facID)->tuple[list[str],list[int]]:
        listCourseNames = []
        listCourseIDs = []
        mycursor.execute("SELECT coursename, courseid from regie.course where professorid =%s",[facID])
        for x in mycursor:
            courseID = x[0]
            coursename = x[1]
            listCourseNames.append(coursename)
            listCourseIDs.append(courseID)
        return listCourseNames, listCourseIDs
    
    def createCourse(courseID,courseName,professorID,bld,roomNum,maxSize,minSize,projector,whiteboard,restricted):
        mycursor.execute("""INSERT INTO regie.course (courseid, coursename, professorid, bld, room, maxsize, 
            minsize, projector, whiteboard, restricted) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s) """,
            [courseID,courseName,professorID,bld,roomNum,maxSize,minSize,projector,whiteboard,restricted])
        db.commit() 


class LogTableInterface:
    #need to determine if sql is mutable to change
    def numRegAndWait(courseid)->int:
        mycursor.execute("SELECT COUNT(*) FROM regie.log WHERE courseid = %s AND action IN ('waiting', 'registered')",[courseid])
        for x in mycursor:
            return x[0]


    def studNumRegCourse(userid)->int:
        mycursor.execute("SELECT COUNT(*) FROM regie.log WHERE userid = %s and action IN ('waiting','registered')",[userid])
        for x in mycursor:
            return x[0]
    
    def studReqWaitCourse(userid)->list[int]:
        listCourseID = []
        mycursor.execute("SELECT courseid FROM regie.log WHERE userid = %s and action IN ('waiting')",[userid])
        for x in mycursor:
            listCourseID.append(x[0])
        return listCourseID
    
    def studReqRegCourse(userid)->list[int]:
        listCourseID = []
        mycursor.execute("SELECT courseid FROM regie.log WHERE userid = %s and action IN ('registered')",[userid])
        for x in mycursor:
            listCourseID.append(x[0])
        return listCourseID

    def hasNotRequest(userid,courseid)->bool: #will call again if returns None then does not exist in table
        mycursor.execute("SELECT ACTION FROM regie.log WHERE userid = %s and courseid = %s",[userid,courseid])
        for x in mycursor:
            if x[0] == "waiting":
                return True
            elif x[0] == None:
                return True
            else:
                return False
    
    def hasRequestWait(userid,courseid)->bool: #will call again if returns None then does not exist in table
        mycursor.execute("SELECT ACTION FROM regie.log WHERE userid = %s and courseid = %s",[userid,courseid])
        for x in mycursor:
            if x[0] == "waiting":
                return True
            else:
                return False
    
    def hasRequestApproved(userid,courseid)->bool:
        mycursor.execute("SELECT ACTION FROM regie.log WHERE userid = %s and courseid = %s",[userid,courseid])
        for x in mycursor:
            if x[0] == "registered":
                return True
            else:
                return False   

    def hasRequestDenied(userid,courseid)->bool:
        mycursor.execute("SELECT ACTION FROM regie.log WHERE userid = %s and courseid = %s",[userid,courseid])
        for x in mycursor:
            if x[0] == "denied":
                return True
            else:
                return False
    
    def submitRestCourseReq(userid,courseid,action): #Restricted Course Request
        mycursor.execute("INSERT INTO regie.log (userid, courseid, action) VALUES (%s,%s,%s)",[userid,courseid,action])
        db.commit()

    def submitRegCourseReq(userid,courseid,action): #Regular Course Request
        mycursor.execute("INSERT INTO regie.log (userid, courseid, action) VALUES (%s,%s,%s)",[userid,courseid,action])
        db.commit()

    def lookUpLogID(userid,courseid)->int:
        mycursor.execute("SELECT logid FROM regie.log WHERE userid = %s and courseid = %s",[userid,courseid])
        for x in mycursor:
            return x[0]

    def updateLogStatus(logid,status):
        mycursor.execute("UPDATE regie.log SET action = %s WHERE logid = %s",[status,logid])
        db.commit()


