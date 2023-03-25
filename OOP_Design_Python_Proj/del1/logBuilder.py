from distutils.command.build import build
from course import Course
import mysql.connector

db = mysql.connector.connect(
    host="localhost",
    user="root",
    passwd="scoutwhitehouse",
    database="regie"
)

mycursor = db.cursor()

class LogFactory:
    @classmethod
    def buildLogEntry(self, numTrans):
        for i in range(numTrans):
            userid = 10000+i #we based on the order that this will work because users with these ids have been created 
            courseid = 56000
            action = 'registered'
            mycursor.execute("""INSERT INTO regie.log (userid, courseid, action) VALUES (%s, %s, %s) """,[userid,courseid,action])
        for i in range(6):
            userid = 10000+i
            mycursor.execute("""INSERT INTO regie.log (userid, courseid, action) VALUES (%s, 57000,'Registered')""",[userid])
    
    def initialize(numOfLog):
        LogFactory.buildLogEntry(numOfLog)
        db.commit()
