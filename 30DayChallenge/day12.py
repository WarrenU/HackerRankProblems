class Person:

	def __init__(self, firstName, lastName, idNumber):
		self.firstName = firstName
		self.lastName = lastName
		self.idNumber = idNumber

	def printPerson(self):
		print("Name:", self.lastName + ",", self.firstName)
		print("ID:", self.idNumber)


class Student(Person):
    #   Class Constructor
    #
    #   Parameters:
    #   firstName - A string denoting the Person's first name.
    #   lastName - A string denoting the Person's last name.
    #   id - An integer denoting the Person's ID number.
    #   scores - An array of integers denoting the Person's test scores.
    #
    # Write your constructor here
    def __init__(self, fn, ln, id, scores_list):
        Person.__init__(self, fn, ln, id)
        self.scores = scores_list

    #   Function Name: calculate
    #   Return: A character denoting the grade.
    #
    # Write your function here
    def calculate(self):
        avg_grade = sum(self.scores)/len(self.scores)
        return (
            'O' if avg_grade >= 90
            else 'E' if avg_grade >=80
            else 'A' if avg_grade >= 70
            else 'P' if avg_grade >= 55
            else 'D' if avg_grade >= 40
            else 'T'
        )
