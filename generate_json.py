import json
import argparse
import random

parser = argparse.ArgumentParser()

parser.add_argument("-n", "--num_entries", 
	type=int, help="Number of json objects in file")

args = parser.parse_args()

NUM_ENTRIES = args.num_entries


names = ['Vedang', 'Mehul', 'Mukesh', 'Akash', 'Paridhi', 'Reshmi', 'Soutreya', 'Sagar', 'Aniket', 'Rishabh', 'Shreshtha', 'Rajkaran']
ids = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]
phones = ['7069911144', '9585731358', '9876543210', '8765432109', '7654321098', '9012345678', '8901234567', '7890123456', '7869504321', '9998887770', '9870654321', '8937474824']
password = ['adsaji2@@', 'fdad24&*', '7&24wrqgeg', 'KJDJ@874wd', 'kfsjaigjd', 'fjdagi3342', 'l301032i0', 'fdw19t2325', 'r2vr2r423', 'eg1e3131ff', 'r11231r31f', 'w1431edsgd']
order_placed = ['Yes', 'No']
order_valid = ['Yes', 'No']


with open('sample.json', 'w') as f:
    f.write('[\n')
    
    for i in range(NUM_ENTRIES):
        randval = random.randint(0, 11)
        
        data = dict()
        
        data['username'] = names[randval]
        data['user_id'] = ids[randval]
        data['phone_no'] = phones[randval]
        data['password'] = password[randval]
        data['price'] = random.randint(1, 200)
        data['num_installments'] = random.randint(1, 10)
        data['product_weight'] = random.randint(2, 50)
        data['order_number'] = i + 1
        data['order_placed'] = order_placed[random.randint(0, 1)]
        data['order_valid'] = order_valid[random.randint(0, 1)]

        json_data = json.dumps(data, indent=4)
        
        if i != NUM_ENTRIES - 1: 
            f.write(json_data + ',\n')
        else:
            f.write(json_data + '\n')
        
    f.write(']')