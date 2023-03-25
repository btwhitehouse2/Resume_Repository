from distutils.command.build import build
import random
from course import Course
import mysql.connector

db = mysql.connector.connect(
    host="localhost",
    user="root",
    passwd="scoutwhitehouse",
    database="regie"
)

mycursor = db.cursor()

class CourseFactory:
    @classmethod
    def buildCourses(self, numTrans):
        for i in range(numTrans-2):
            courseID = 51100+i 
            #since faculty have already been built will follow same procedure for assigning ID as they did in userFactory
            professorID = 40000+i 
            bld = random.choice(["JCL","Ryerson"])
            roomNum = random.randint(100,399)
            maxSize = random.randint(20,30)
            projector = random.choice([True,False])
            whiteboard = random.choice([True,False])
            courseName = ("Unix Systems","Object Oriented Programming","Block Chain","Compilers",
            "Parallel Programming","Product Management","Python Programming","Web Development")[i]
            minSize = 5
            courseID = 51100+i #courseNum = random.randint(51100,56999)
            restricted = False
            mycursor.execute("""INSERT INTO regie.course (courseid, coursename, professorid, bld, room, maxsize, 
            minsize, projector, whiteboard, restricted) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s) """,
            [courseID,courseName,professorID,bld,roomNum,maxSize,minSize,projector,whiteboard,restricted])
        mycursor.execute("""INSERT INTO regie.course (courseid, coursename, professorid, bld, room, maxsize, 
        minsize, projector, whiteboard, restricted) VALUES (56000, 'Algorithms', 42345,'Ryerson', 150, 25, 5, true, true, 
        true) """)
        mycursor.execute("""INSERT INTO regie.course (courseid, coursename, professorid, bld, room, maxsize, 
        minsize, projector, whiteboard, restricted) VALUES (57000, 'Discrete Math', 42345,'Ryerson', 150, 6, 5, true, true, 
        false) """)

    def initialize(numCourse):
        CourseFactory.buildCourses(numCourse)
        db.commit()
