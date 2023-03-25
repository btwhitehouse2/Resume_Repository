from abc import ABC
from ast import Pass
from interface import UserInterface, UserTableInterface, CourseTableInterface

class User(ABC):
    def __init__(self,ID,firstName,lastName,userName,email,_password):
        self.ID = ID
        self.firstName = firstName
        self.lastName = lastName
        self.userName = userName
        self._password = _password
        self.email = email

    def viewProfile(self)->str:
        id = self.ID
        return UserTableInterface.returnUserInfo(id)
    
    def __str__(self):
        return f"""
        Full Name: {self.firstName} {self.lastName}
        UChicago User Name: {self.userName}
        UChicago Email: {self.email}
    """

class Faculty(User):
    def __init__(self,ID,firstName,lastName,email,userName,_password):
        super().__init__(ID,firstName,lastName,email,userName,_password)
        self.role = "faculty"

    def acceptStud(self,courseID,studentID)->str:
        acceptStat = True
        return UserInterface.RecieveFacultyResponse(courseID,studentID,self.ID,acceptStat,self.firstName)

    def declStud(self,courseID,studentID)->str:
        declStat = False
        return UserInterface.RecieveFacultyResponse(courseID,studentID,self.ID,declStat,self.firstName)

    def viewCourses(self)->str:
        return UserInterface.StatusFacCourse(self.ID,self.firstName)

    def reqNewCourse(self,courseName,AdminId,restricted)->str:
        return UserInterface.NotifyAdminOfProfRequest(courseName,self.ID,restricted,self.firstName,self.lastName,AdminId)

class FullTimeFaculty(Faculty):
    def __init__(self,ID,firstName,lastName,email,userName,_password):
        super().__init__(ID,firstName,lastName,email,userName,_password)
        self.status = "Full Time"

    def __str__(self):
        return f"""
        Full Name: {self.firstName} {self.lastName}
        User ID: {self.ID}
        Role + Status: {self.role} {self.status}
        UChicago User Name: {self.userName}
        UChicago Email: {self.email}
    """

class Student(User):
    def __init__(self,ID,firstName,lastName,email,userName,_password):
        super().__init__(ID,firstName,lastName,email,userName,_password)
        self.role = "student"

    def registerForCourse(self, courseNum)->str:
        return UserInterface.recieveCourseRequest(self.ID,courseNum,self.firstName,self.lastName)

    def checkStatus(self)->str:
        return UserInterface.sendCurrentCourseRegistered(self.ID,self.firstName)

class FullTimeStudent(Student):
    def __init__(self,ID,firstName,lastName,email,userName,_password):
        super().__init__(ID,firstName,lastName,email,userName,_password)
        self.status = "Full Time"
    
    def __str__(self):
        return f"""
        Full Name: {self.firstName} {self.lastName}
        Student ID: {self.ID}
        Role + Status: {self.role} {self.status}
        UChicago User Name: {self.userName}
        UChicago Email: {self.email}
    """

class Administrator(User):
    def __init__(self,ID,firstName,lastName,email,userName,_password): 
        super().__init__(ID,firstName,lastName,email,userName,_password)
        self.role = "Administrator"

    def __str__(self):
        return f"""
        Full Name: {self.firstName} {self.lastName}
        Admin ID: {self.ID}
        Role: {self.role}
        UChicago User Name: {self.userName}
        UChicago Email: {self.email}
    """

    def openRegistration(self)->str:
        return UserInterface.OpenRegistration(self.ID,self.firstName,self.lastName)
    
    def addCourse(self,courseName,courseID,ProfID,bld,room,mxSize,mnSize,proj,wb,rest)->str:
        return UserInterface.SubmitNewCourse(self.ID,self.firstName,self.lastName,courseID,courseName,ProfID,bld,room,mxSize,mnSize,proj,wb,rest)
