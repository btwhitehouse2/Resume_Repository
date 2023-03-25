#Execute these commands from MYSQLWorkBench or comprable Development Environment

# CREATE SCHEMA IF NOT EXISTS regie;
# CREATE TABLE IF NOT EXISTS regie.user ( 
# 	userid int PRIMARY KEY, userfirst VARCHAR(12), 
# 	userlast VARCHAR(12), email VARCHAR(30), 
# 	username VARCHAR(12), password VARCHAR(20), 
# 	role VARCHAR(12), status VARCHAR(12));
# CREATE TABLE IF NOT EXISTS regie.course (
# 	courseid int PRIMARY KEY, coursename VARCHAR(12),
#     professorid int, bld VARCHAR(12), 
#     room int, maxsize int, 
#     minsize int, projector bool, 
#     whiteboard bool, restricted bool,
#     FOREIGN KEY (professorid) REFERENCES regie.user(userid)
#     );
# CREATE TABLE IF NOT EXISTS regie.log (
#   logid INT AUTO_INCREMENT PRIMARY KEY, userid INT NOT NULL,
#   courseid INT NOT NULL, action VARCHAR(12) NOT NULL,
#   FOREIGN KEY (userid) REFERENCES regie.user(userid),
#   FOREIGN KEY (courseid) REFERENCES regie.course(courseid)
# );