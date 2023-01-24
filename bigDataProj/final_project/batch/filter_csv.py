import pandas as pd, sys, re

def filt_csv(input_csv,output_csv):
    filtered_csv = []
    # reader = (open('/Users/btwhi/Documents/Big_Data/final_project/Cook_County_Foreclosure_2013_through_March_27_2015.csv', 'r'))
    # writer = (open('/Users/btwhi/Documents/Big_Data/final_project/batch/output.csv', 'w'))
    reader = (open(input_csv, 'r'))
    writer = (open(output_csv, 'w'))
    for row in reader.readlines():
        row = row.rstrip()
        firstChar = [char for char in str(row)][0]
        int_list = ['0','1','2','3','4','5','6','7','8','9']
        if firstChar not in int_list:
            try:
                last_elm = str(filtered_csv.pop()).rstrip()
                filtered_csv.append(last_elm + row +"\n")
            except:
                filtered_csv.append(row + "\n")
        else:
            filtered_csv.append(row + "\n")

    for line in filtered_csv:
        #use 2 commas in a row to determine if there is a pieces of missing data exclude it from the ground truth
        if ",," in line:
            next
        elif 'PIN' in line:
            writer.write(line)
        elif '00000-0000' in line:
            next 
        elif 'LIS PENDENS FORECLOSURE' in line:
            next
        elif 'QUIT CLAIM DEED' in line:
            next
        elif line == '':
            next
        else:
            try:
                line8 = re.search(r"(\d{5})([-]\d{4})", line)             
                line8.groups()
                zipFull = line8.group()
                zipFirstFive =line8.group(1)
                filterline = line.replace(zipFull,zipFirstFive,1)
                writer.write(filterline)
            except TypeError and AttributeError:
                next


if __name__ == "__main__":
    # input_csv = sys.argv[0]
    # output_csv = sys.argv[1]
    filt_csv('/Users/btwhi/Documents/Big_Data/final_project/batch/Cook_County_Recorder_2013_through_March_27__2015.csv','/Users/btwhi/Documents/Big_Data/final_project/batch/Filtered_CC_Mortgage_2013_through_March_27_2015.csv')
    #filt_csv('/Users/btwhi/Documents/Big_Data/final_project/batch/testIn.csv','/Users/btwhi/Documents/Big_Data/final_project/batch/testOut.csv')
    # filt_csv(input_csv,output_csv)
    # reader2 = pd.read_csv(output_csv)
    # reader2 = pd.read_csv('/Users/btwhi/Documents/Big_Data/final_project/batch/Filtered_Cook_County_Foreclosure_2013_through_March_27_2015.csv')
    # print("sum totals for coumns with blank fields",reader2.isnull().sum())