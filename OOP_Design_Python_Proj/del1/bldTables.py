import mysql.connector

#Execute these commands from MYSQLWorkBench or comprable Development Environment

def initialize():
    db = mysql.connector.connect(
        host="localhost",
        user="root",
        passwd="scoutwhitehouse",
        database="regie"
    )

    #drop log table
    stat1 = db.cursor()
    #drop course table
    stat2 = db.cursor() 
    #drop user table
    stat3 = db.cursor()
    #build log table
    stat4 = db.cursor()
    #build course table
    stat5 = db.cursor()
    #build user table
    stat6 = db.cursor()

    stat1.execute("""DROP TABLE regie.log""")
    stat2.execute("""DROP TABLE regie.course""")
    stat3.execute("""DROP TABLE regie.user""")
    stat4.execute("""CREATE TABLE IF NOT EXISTS regie.user ( 
        userid int PRIMARY KEY, userfirst VARCHAR(12), 
        userlast VARCHAR(12), email VARCHAR(40), 
        username VARCHAR(20), password VARCHAR(20), 
        role VARCHAR(12), status VARCHAR(12))""")
    stat5.execute("""CREATE TABLE IF NOT EXISTS regie.course (courseid int 
        PRIMARY KEY, coursename VARCHAR(30),
        professorid int, bld VARCHAR(12), 
        room int, maxsize int, 
        minsize int, projector bool, 
        whiteboard bool, restricted bool,
        FOREIGN KEY (professorid) REFERENCES regie.user(userid)
        )""")
    stat6.execute("""CREATE TABLE IF NOT EXISTS regie.log (
        logid INT AUTO_INCREMENT PRIMARY KEY, userid INT NOT NULL,
        courseid INT NOT NULL, action VARCHAR(12) NOT NULL,
        FOREIGN KEY (userid) REFERENCES regie.user(userid),
        FOREIGN KEY (courseid) REFERENCES regie.course(courseid))
        """)
    #commits to database
    db.commit()
