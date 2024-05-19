import pandas as pd
import random
import string
from faker import Faker

fake = Faker()

# Function to generate a random student ID
def generate_student_id():
    return ''.join(random.choices(string.ascii_uppercase + string.digits, k=8))

# Function to generate a random class
def generate_class():
    return random.choice(['Freshman', 'Sophomore', 'Junior', 'Senior'])

# Function to generate a random email based on the name
def generate_email(name):
    return name.replace(" ", ".").lower() + "@university.edu"

# Lists to hold the generated data
student_ids = []
names = []
classes = []
emails = []
addresses = []

for _ in range(1000):
    name = fake.name()
    student_ids.append(generate_student_id())
    names.append(name)
    classes.append(generate_class())
    emails.append(generate_email(name))
    addresses.append(fake.address().replace("\n", ", "))

# Creating a DataFrame
data = {
    "Student ID": student_ids,
    "Name": names,
    "Class": classes,
    "Email": emails,
    "Address": addresses
}

df = pd.DataFrame(data)

# Saving the DataFrame to an Excel file
df.to_excel("student_records.xlsx", index=False)

print("Excel file 'student_records.xlsx' has been created.")
