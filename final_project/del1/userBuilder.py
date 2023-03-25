import random
import mysql.connector
from user import FullTimeStudent, FullTimeFaculty, Administrator

db = mysql.connector.connect(
    host="localhost",
    user="root",
    passwd="scoutwhitehouse",
    database="regie"
)

mycursor = db.cursor()

listFirstName = ["John","Jane","Jim","Jannet","James","Robert","Michael","Beth",
"Brian","Keri","Kevin","Mary","Patricia","Jennifer","Linda","Elizabeth","Barbara","Susan",
"Jessica","Sarah","Karen","Lisa","Nancy","Betty","Margaret","Sandra","Ashley","Emily","Donna",
"David","William","Richard","Joseph","Thomas","Charles","Christopher","Daniel","Matthew","Anthony"]

listLastName = ["Smith","Johnson","Kim","Zhang","White","Martin","Jones","Williams","Brown","Garcia",
"Miller","Davis","Rodriguez","Martinez","Hernandez","Lopez","Gonzales","Wilson","Anderson","Thomas",
"Taylor","Moore","Jackson","Martin","Lee","Perez","Thompson","Harris","Sanchez","Clark","Ramirez"
,"Lewis","Robinson"]

class UserFactory:
    @classmethod
    def buildFullTimeStudents(self, numTrans):
        for i in range(numTrans-1):
            id = 10000+i
            firstName = random.choice(listFirstName)
            lastName = random.choice(listLastName)
            userName = firstName + lastName
            password = "password1234!"
            email = userName+"@uchicago.edu"
            role = "student"
            status = "full"
            mycursor.execute("INSERT INTO regie.user (userid, userfirst, userlast, email, username, password, role, status) VALUES (%s,%s,%s,%s,%s,%s,%s,%s) ",[id,firstName,lastName,email,userName,password,role,status])
        mycursor.execute("INSERT INTO regie.user (userid, userfirst, userlast, email, username, password, role, status) VALUES (12276,'Brian','Whitehouse','bwhitehouse@uchicago.edu','bwhitehouse','password','student','full') ")
        bWhitehouse = FullTimeStudent(12276,"Brian","Whitehouse","bwhitehouse@uchicago.edu","bwhitehouse","password")
        return bWhitehouse

    @classmethod
    def buildFullTimeFactulty(self, numTrans):
        for i in range(numTrans-1):
            id = 40000+i #streamlining this in order to assign faculty to courses
            firstName = random.choice(["Geraldine","Michael","Lamont","Shelley"])
            lastName = random.choice(["Brady","Spertus","Samuels","Rossell"])
            userName = firstName + lastName
            password = "password1234!"
            email = userName+"@uchicago.edu"
            role = "faculty"
            status = "full"
            mycursor.execute("INSERT INTO regie.user (userid, userfirst, userlast, email, username, password, role, status) VALUES (%s,%s,%s,%s,%s,%s,%s,%s) ",[id,firstName,lastName,email,userName,password,role,status])
        mycursor.execute("INSERT INTO regie.user (userid, userfirst, userlast, email, username, password, role, status) VALUES (42345,'Tim','Ng','tng@uchicago.edu','tng','password','faculty','full') ")
        tNg = FullTimeFaculty(42345,"Tim","Ng","tng@uchicago.edu","tng","password")
        return tNg

    @classmethod
    def buildAdministrator(self, numTrans)->Administrator:
        for i in range(numTrans-1):
            id = random.randint(90000,99998)
            firstName = random.choice(["Molly","Karin","Megan"])
            lastName = random.choice(["Stoner","Czaplewski","Woodward"])
            userName = firstName + lastName
            password = "password1234!"
            email = userName+"@uchicago.edu"
            role = "admin"
            status = ""
            mycursor.execute("INSERT INTO regie.user (userid, userfirst, userlast, email, username, password, role, status) VALUES (%s,%s,%s,%s,%s,%s,%s,%s) ",[id,firstName,lastName,email,userName,password,role,status])
        mycursor.execute("INSERT INTO regie.user (userid, userfirst, userlast, email, username, password, role, status) VALUES (99999,'Bahareh','Lampert','blampert@uchicago.edu','baharehlampert','password','admin','') ")
        bLampert = Administrator(99999,"Bahareh","Lampert","blampert@uchicago.edu","baharehlampert","password")
        return bLampert
        
    def initialize(numAdmin,NumFac,NumStud):
        bLampert = UserFactory.buildAdministrator(numAdmin)
        tNg = UserFactory.buildFullTimeFactulty(NumFac)
        bWhitehouse = UserFactory.buildFullTimeStudents(NumStud)
        db.commit()
        return bLampert, tNg, bWhitehouse


