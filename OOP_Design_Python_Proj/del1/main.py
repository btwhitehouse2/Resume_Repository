import userBuilder as UF
import courseBuilder as CF
import logBuilder as LF

def main():
    #Initiate tables in a Database
    #I accomplish this using MySQL WorkBench and the commands saved in bldTables.py
    #Next Initiate the data by using User Factory (fulltime student, faculty, administrator)
    #Since User ID is referenced in Course need to initiate this first
    uf = UF.UserFactory
    numberAdmin=2;numberFac=5;NumberStud=20
    admin, fullTimefac, fullTimestud = uf.initialize(numberAdmin,numberFac,NumberStud)
    cf = CF.CourseFactory
    #need to ensure that there are at least 1 more faculty than course for the factory to work
    cf.initialize(numCourse=3)
    lf = LF.LogFactory
    lf.initialize(numOfLog=5)
    #Will print each object to demonstrate use of inheritance and classes
    print(admin)
    print(fullTimefac)
    print(fullTimestud)
    #These objects also exist in the database will query to demonstrate there concurent existance
    print("--------From query of table------")
    print(*admin.viewProfile())
    print(*fullTimefac.viewProfile())
    print(*fullTimestud.viewProfile())
    print('\n')
    print("--------Scenario 1 Registering for a Restricted Class------")
    #Course request is for Algorithms
    print(fullTimestud.registerForCourse(56000))
    #Oh not course registration isn't open
    print(admin.openRegistration())
    #check the status of the student for courses
    print(fullTimestud.checkStatus())
    #Now try to register for a course again
    print(*fullTimestud.registerForCourse(56000))
    print('\n')
    #Now Check the registration Status
    print(fullTimestud.checkStatus())
    #Faculty member recieves the request and now approves the request
    #Must supply the courseID and studentID 
    print(fullTimefac.acceptStud(56000,12276))
    #NowCheckStatusAgain
    print(fullTimestud.checkStatus())
    print('\n')
    print("--------Scenario 2 Registering for a Class that is full ------")
    #Full Time Faculty checks on the status of classes that they are teaching
    print(fullTimefac.viewCourses())
    #One course is full so the faculty 
    print(fullTimestud.registerForCourse(57000))
    print('\n')
    print("--------Scenario 3 Faculty requests creation of new course from Admin ------")
    #Faculty member needs to request a new course with the admin
    #Faculty will need to supply the course name the admin ID and if the course is restricted
                                    #name             #Admin ID   #Is Restricted
    print(fullTimefac.reqNewCourse("Researching Algorithms",99999,False))
    #Admin needs to add a new course to the DB
    #Admin will have to select a lot of the adminstrative information ie. location and size
                            #course name         #courID #ProfID#BLD #RM #Size  #WB #Proj #Restricted
    print(admin.addCourse("Researching Algorithms",58000,42345,"JCL",203,20,5,True,True,False))
    #Now the faculty can view and see that they are teaching another course
    print(fullTimefac.viewCourses())


if __name__ == "__main__":
    main()