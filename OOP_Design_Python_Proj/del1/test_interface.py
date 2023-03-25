import unittest

from interface import UserInterface, UserTableInterface, CourseTableInterface, LogTableInterface

class TestLogTableInterface(unittest.TestCase):
    """test some of the interface for the Log Table"""

    def setUp(self):
        ##Will build a factor
        """utilizes a set up function in order to reduce having to build a system"""
        self.RegistrationSpace = {}
    
    def test_numRegAndWait(self):
        """ testing the number of students registered and waiting for confirmation
        for a specific course """
        courseid = 56000
        result = LogTableInterface.numRegAndWait(courseid)
        self.assertEqual(result,5)

    def test_studNumRegCourse(self):
        """ testing the number of classes that a specific student is registered and 
        or waiting for confirmation """
        userid = 10001
        result = LogTableInterface.studNumRegCourse(userid)
        self.assertEqual(result,1)

class TestCourseTableInterface(unittest.TestCase):
    """ test the interface for the course table"""
    
    def test_isRestricted(self):
        """Checks to see if a course is listed as restricted"""
        courseid = 56000
        result = CourseTableInterface.isRestricted(courseid)
        self.assertTrue(result)

    def test_isFull(self):
        """Checks to see if a course is full"""
        courseid = 56000
        result = CourseTableInterface.isFull(courseid)
        self.assertFalse(result)
    
    def test_maxSize(self):
        """Confirms the max size of a course"""
        courseid = 56000
        result = CourseTableInterface.maxSize(courseid)
        self.assertEqual(result,6)

    def test_returnCourseInfo(self):
        """Confirms the max size of a course"""
        courseid = 56000
        result = CourseTableInterface.returnCourseInfo(courseid)
        professorID = 42345
        course = "Algorithms"
        self.assertEqual(result,(course,professorID))

class TestUserTableInterface(unittest.TestCase):
    """test the interface for the user table"""
    
    def test_isStudent(self):
        """Checks to see if the user is a student"""
        userid = 10001
        result = UserTableInterface.isStudent(userid)
        self.assertTrue(result)
    
    def test_isAdmin(self):
        """Checks to see if the user is an admin, will intentionally put a student
        id to see if it will return false"""
        userid = 10000
        result = UserTableInterface.isAdmin(userid)
        self.assertFalse(result)
    
    def test_isFullTime(self):
        """Checks the status of the user to see if they are full time"""
        userid = 10001
        result = UserTableInterface.isFullTime(userid)
        self.assertTrue(result)
    
    def test_hasRoomInSchedule(self):
        """Checks to see if the user has room in schedule to take more classes"""
        userid = 10002
        result = UserTableInterface.hasRoomInSchedule(userid)
        self.assertTrue(result)

class TestUserInterface(unittest.TestCase):
    """Test the User Interface class"""
    
    def test_OpenRegistration(self):
        """Confirms that the registration can be opened by an admin"""
        u = UserInterface(False)
        userid = 99999
        u.OpenRegistration(userid)
        self.assertTrue(u.isRegOpen)

    def test_CloseRegistration(self):
        """Confirms that the registration can be closed by an admin"""
        u = UserInterface(True)
        userid = 99999
        u.CloseRegistration(userid)
        self.assertFalse(u.isRegOpen)

    def test_recieveCourseRequest(self):
        pass    

if __name__ == '__main__':
    unittest.main()
