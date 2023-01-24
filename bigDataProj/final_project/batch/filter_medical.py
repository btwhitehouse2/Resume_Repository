import pandas as pd, sys, re

def filt_csv(input_csv,output_csv):
    # reader = (open('/Users/btwhi/Documents/Big_Data/final_project/Cook_County_Foreclosure_2013_through_March_27_2015.csv', 'r'))
    # writer = (open('/Users/btwhi/Documents/Big_Data/final_project/batch/output.csv', 'w'))
    reader = (open(input_csv, 'r'))
    writer = (open(output_csv, 'w'))
    for line in reader.readlines():
        #use 2 semi colons in a row to determine if there is a pieces of missing data exclude it from the ground truth
        line = line.rstrip()
        if "PENDING" in line:
            next
        # elif ';;' in line:
        #     next
        elif 'Case Number' in line:
            line = "ACC;HOM;NAT;SUC;UND;"+ line +"\n"          
            writer.write(line) 
        elif line == '':
            next
        else:
            if "ACCIDENT;" in line:
                line = "1;0;0;0;0;"+ line +"\n"
                writer.write(line)
            elif "HOMICIDE;" in line:
                line = "0;1;0;0;0;"+ line +"\n"
                writer.write(line)
            elif "NATURAL;" in line:
                line = "0;0;1;0;0;" + line + "\n"
                writer.write(line)
            elif "SUICIDE;" in line:
                line = "0;0;0;1;0;" + line + "\n"
                writer.write(line)
            elif "UNDETERMINED;" in line:
                line = "0;0;0;0;1;" + line + "\n"
                writer.write(line)
            else:
                next


if __name__ == "__main__":
    # input_csv = sys.argv[0]
    # output_csv = sys.argv[1]
    filt_csv('/Users/btwhi/Documents/Big_Data/final_project/batch/Medical_Examiner_Case_Archive.csv','/Users/btwhi/Documents/Big_Data/final_project/batch/Filtered_Medical_Examiner.csv')
    #filt_csv('/Users/btwhi/Documents/Big_Data/final_project/batch/testIn.csv','/Users/btwhi/Documents/Big_Data/final_project/batch/testOut.csv')
    # filt_csv(input_csv,output_csv)
    # reader2 = pd.read_csv(output_csv)
    # reader2 = pd.read_csv('/Users/btwhi/Documents/Big_Data/final_project/batch/Filtered_Cook_County_Foreclosure_2013_through_March_27_2015.csv')
    # print("sum totals for coumns with blank fields",reader2.isnull().sum())