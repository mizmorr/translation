import json
import re
import os

CLASSES_OF_TOKENS = ['W', 'I', 'O', 'R', 'N', 'C']

def lab1():

    SERVICE_WORDS = ['abstract', 'case', 'continue', 'extends', 'goto', 'int', 'package', 'short',
                    'try', 'assert', 'catch', 'default', 'final', 'if', 'private',
                    'static', 'this', 'void', 'boolean', 'char', 'do','long', 'protected',
                    'throw', 'volatile', 'break', 'class', 'double', 'float', 'import', 'native',
                    'public', 'super','throws','while','byte','const','else','for','instanceof',
                    'new','return','switch','transient','print','println','main','System',
                    'out','String','args','in.nextInt()']

    OPERATIONS = ['*','+','-','%', '/','++','*=','+=','-=','%=','/=','==', '<', '<=', '!=', '=', '>', '>=','&','|']

    SEPARATORS = ['\t', '\n', ' ', '(', ')', ',', '.', ':', ';', '[', ']','{','}']


    def check(tokens, token_class, token_value):
        if not(token_value in tokens[token_class]):
            token_code = str(len(tokens[token_class]) + 1)
            tokens[token_class][token_value] = token_class + token_code

    def get_operation(input_sequence, i):
        for k in range(2, 0, -1):
            if i + k < len(input_sequence):
                buffer = input_sequence[i:i + k]
                if buffer in OPERATIONS:
                    return buffer
        return ''

    def get_separator(input_sequence, i):
        buffer = input_sequence[i]
        if buffer in SEPARATORS:
            return buffer
        return ''

    # лексемы
    tokens = {'W': {}, 'I': {}, 'O': {}, 'R': {}, 'N': {}, 'C': {}}

    for service_word in SERVICE_WORDS:
        check(tokens, 'W', service_word)
    for operation in OPERATIONS:
        check(tokens, 'O', operation)
    for separator in SEPARATORS:
        check(tokens, 'R', separator)

    # файл, содержащий текст на входном языке программирования
    f = open('script/java.txt', 'r')
    input_sequence = f.read()
    input_sequence = re.sub(r"(\w)\+\+", r"\1 = \1 + 1", input_sequence)
    input_sequence = input_sequence.replace("public static void main(String[] args)\n{\n","\n\n")
    input_sequence = input_sequence.replace("\n}\n","")
    input_sequence = input_sequence.replace("public static mult (int a, int b) {","\n\n")
    input_sequence = input_sequence.replace("System.out.println(","alert(")
    f.close()

    i = 0
    state = 'S'
    output_sequence = buffer = ''
    while i < len(input_sequence):
        symbol = input_sequence[i]
        operation = get_operation(input_sequence, i)
        separator = get_separator(input_sequence, i)
        if state == 'S':
            buffer = ''
            if symbol.isalpha():
                state = 'q1'
                buffer += symbol
            elif symbol.isdigit():
                state = 'q3'
                buffer += symbol
            elif symbol == "'":
                state = 'q9'
                buffer += symbol
            elif symbol == '"':
                state = 'q10'
                buffer += symbol
            elif symbol == '/':
                state = 'q11'
            elif operation:
                check(tokens, 'O', operation)
                output_sequence += tokens['O'][operation] + ' '
                i += len(operation) - 1
            elif separator:
                if separator != ' ':
                    check(tokens, 'R', separator)
                    output_sequence += tokens['R'][separator]
                    if separator == '\n':
                        output_sequence += '\n'
                    else:
                        output_sequence += ' '
            elif i == len(input_sequence) - 1:
                state = 'Z'
        elif state == 'q1':
            if symbol.isalpha():
                buffer += symbol
            elif symbol.isdigit():
                state = 'q2'
                buffer += symbol
            else:
                if operation or separator:
                    if buffer in SERVICE_WORDS:
                        output_sequence += tokens['W'][buffer] + ' '
                    elif buffer in OPERATIONS:
                        output_sequence += tokens['O'][buffer] + ' '
                    else:
                        check(tokens, 'I', buffer)
                        output_sequence += tokens['I'][buffer] + ' '
                    if operation:
                        check(tokens, 'O', operation)
                        output_sequence += tokens['O'][operation] + ' '
                        i += len(operation) - 1
                    if separator:
                        if separator != ' ':
                            check(tokens, 'R', separator)
                            output_sequence += tokens['R'][separator]
                            if separator == '\n':
                                output_sequence += '\n'
                            else:
                                output_sequence += ' '
                state = 'S'
        elif state == 'q2':
            if symbol.isalnum():
                buffer += symbol
            else:
                if operation or separator:
                    check(tokens, 'I', buffer)
                    output_sequence += tokens['I'][buffer] + ' '
                    if operation:
                        check(tokens, 'O', operation)
                        output_sequence += tokens['O'][operation] + ' '
                        i += len(operation) - 1
                    if separator:
                        if separator != ' ':
                            check(tokens, 'R', separator)
                            output_sequence += tokens['R'][separator]
                            if separator == '\n':
                                output_sequence += '\n'
                            else:
                                output_sequence += ' '
                    state = 'S'
        elif state == 'q3':
            if symbol.isdigit():
                buffer += symbol
            elif symbol == '.':
                state = 'q4'
                buffer += symbol
            elif symbol == 'e' or symbol == 'E':
                state = 'q6'
                buffer += symbol
            else:
                if operation or separator:
                    check(tokens, 'N', buffer)
                    output_sequence += tokens['N'][buffer] + ' '
                    if operation:
                        check(tokens, 'O', operation)
                        output_sequence += tokens['O'][operation] + ' '
                        i += len(operation) - 1
                    if separator:
                        if separator != ' ':
                            check(tokens, 'R', separator)
                            output_sequence += tokens['R'][separator]
                            if separator == '\n':
                                output_sequence += '\n'
                            else:
                                output_sequence += ' '
                    state = 'S'
        elif state == 'q4':
            if symbol.isdigit():
                state = 'q5'
                buffer += symbol
        elif state == 'q5':
            if symbol.isdigit():
                buffer += symbol
            elif symbol == 'e' or symbol == 'E':
                state = 'q6'
                buffer += symbol
            else:
                if operation or separator:
                    check(tokens, 'N', buffer)
                    output_sequence += tokens['N'][buffer] + ' '
                    if operation:
                        check(tokens, 'O', operation)
                        output_sequence += tokens['O'][operation] + ' '
                        i += len(operation) - 1
                    if separator:
                        if separator != ' ':
                            check(tokens, 'R', separator)
                            output_sequence += tokens['R'][separator]
                            if separator == '\n':
                                output_sequence += '\n'
                            else:
                                output_sequence += ' '
                    state = 'S'
        elif state == 'q6':
            if symbol == '-' or symbol == '+':
                state = 'q7'
                buffer += symbol
            elif symbol.isdigit():
                state = 'q8'
                buffer += symbol
        elif state == 'q7':
            if symbol.isdigit():
                state = 'q8'
                buffer += symbol
        elif state == 'q8':
            if symbol.isdigit():
                buffer += symbol
            else:
                if operation or separator:
                    check(tokens, 'N', buffer)
                    output_sequence += tokens['N'][buffer] + ' '
                    if operation:
                        check(tokens, 'O', operation)
                        output_sequence += tokens['O'][operation] + ' '
                        i += len(operation) - 1
                    if separator:
                        if separator != ' ':
                            check(tokens, 'R', separator)
                            output_sequence += tokens['R'][separator]
                            if separator == '\n':
                                output_sequence += '\n'
                            else:
                                output_sequence += ' '
                state = 'S'
        elif state == 'q9':
            if symbol != "'":
                buffer += symbol
            elif symbol == "'":
                buffer += symbol
                check(tokens, 'C', buffer)
                output_sequence += tokens['C'][buffer] + ' '
                state = 'S'
        elif state == 'q10':
            if symbol != '"':
                buffer += symbol
            elif symbol == '"':
                buffer += symbol
                check(tokens, 'C', buffer)
                output_sequence += tokens['C'][buffer] + ' '
                state = 'S'
        elif state == 'q11':
            if symbol == '/':
                state = 'q12'
            elif symbol == '*':
                state = 'q13'
        elif state == 'q12':
            if symbol == '\n':
                state = 'S'
            elif i == len(input_sequence) - 1:
                state = 'Z'
        elif state == 'q13':
            if symbol == '*':
                state = 'q14'
        elif state == 'q14':
            if symbol == '/':
                state = 'q15'
        elif state == 'q15':
            if symbol == '\n':
                state = 'S'
            elif i == len(input_sequence) - 1:
                state = 'Z'
        i += 1

    # файлы, содержащие все таблицы лексем
    for token_class in tokens.keys():
        with open('script/'+'%s.json' % token_class, 'w') as write_file:
            data = {val: key for key, val in tokens[token_class].items()}
            json.dump(data, write_file, indent=4, ensure_ascii=False)

    # файл, содержащий последовательность кодов лексем входной программы
    f = open('script/tokens.txt', 'w')
    f.write(output_sequence)
    f.close()

def syntax_analyzer():
    global i
    global nxtsymb
    global row_counter
    i = -1
    nxtsymb = None
    row_counter = 1

    # лексемы
    tokens = {'W': {}, 'I': {}, 'O': {}, 'R': {}, 'N': {}, 'C': {}}

    # файлы, содержащие все таблицы лексем
    for token_class in tokens.keys():
        with open("script/"+'%s.json' % token_class, 'r') as read_file:
            data = json.load(read_file)
            tokens[token_class] = data

    # файл, содержащий последовательность кодов лексем входной программы
    f = open('script/tokens.txt', 'r')
    input_sequence = f.read()
    f.close()

    regexp = '[' + '|'.join(tokens.keys()) + ']' + '\d+'
    match = re.findall(regexp, input_sequence)

    # обработка ошибочной ситуации
    def error():
        out_sq = 'Проверьте правильность кода. Обнаружена синтаксическая ошибка. Строка: '

        f = open('script/error.txt','w')
        out_sq += str(row_counter)

        f.write(out_sq)
        f.close()
        #print('Ошибка в строке', row_counter)
        return


    # помещение очередного символа в nxtsymb
    def scan():
        global i, nxtsymb, row_counter
        i += 1
        if i >= len(match):
            if not(nxtsymb in ['\n', ';', '}']):
                error()
        else:
            for token_class in tokens.keys():
                if match[i] in tokens[token_class]:
                    nxtsymb = tokens[token_class][match[i]]
            if nxtsymb == '\n':
                row_counter += 1
                scan()
            #print(i, row_counter, nxtsymb)

    # программа
    def program():
        operators()

    # операторы
    def operators():
        global i
        scan()
        while name() or \
            nxtsymb in ['int', 'double', 'string', 'boolean', 'float', '{', 'public static void main(String[] args)',\
                        'if', 'for', 'while', 'break', 'continue', 'return']:
            operator()
            if nxtsymb == ';':
                scan()
            if nxtsymb == '}':
                break

    # оператор
    def operator():
        if nxtsymb in ['int','double','float','boolean','string']:
            description()
            if nxtsymb != ';': error()
        elif name():
            scan()
            if nxtsymb == ':':
                scan()
                operator()
            elif nxtsymb == '(':
                scan()
                if nxtsymb != ')':
                    expression()
                    while nxtsymb == ',':
                        scan()
                        expression()
                    if nxtsymb != ')': error()
                scan()
            elif nxtsymb == '=':
                scan()
                expression()
                if nxtsymb != ';': error()
            else: error()
        elif nxtsymb == '{': compound_operator()
        elif nxtsymb == 'public static': function()
        elif nxtsymb == 'if': conditional_operator()
        elif nxtsymb == 'for': for_loop()
        elif nxtsymb == 'while': while_loop()
        elif nxtsymb == 'break':
            break_operator()
            scan()
            if nxtsymb != ';': error()
        elif nxtsymb == 'continue':
            continue_operator()
            scan()
            if nxtsymb != ';': error()
        elif nxtsymb == 'return':
            return_operator()
            if nxtsymb != ';': error()
        else: error()

    # имя (идентификатор)
    def name():
        return nxtsymb in tokens['I'].values() or \
            nxtsymb in ['System.out.println', 'alert']

    # описание
    def description():
        scan()
        if not(name()): error()
        scan()
        if nxtsymb == ',':
            while nxtsymb == ',':
                scan()
                if not(name()): error()
                scan()
        elif nxtsymb == '=':
            scan()
            if nxtsymb == 'new':
                scan()
                if not(name()): error()
                if not(nxtsymb == '['): error()
                scan()
                if not(integer()): error()
                scan()
                if not(nxtsymb == ']'): error()
                scan()
            else:
                expression()

    # список имен
    def list_of_names():
        if not(name()): error()
        scan()
        while nxtsymb == ',':
            scan()
            if not(name()): error()
            scan()

    # функция
    def function():
        if nxtsymb != 'public static void main(String[] args)\n{': error()
        scan()
        if not(name()): error()
        scan()
        if nxtsymb == '(':
            scan()
            if name():
                list_of_names()
        if nxtsymb != ')': error()
        scan()
        compound_operator()

    # выражение
    def expression():
        if nxtsymb == '(':
            scan()
            expression()
            if nxtsymb != ')': error()
            scan()
        elif name():
            scan()
            if nxtsymb == '(':
                scan()
                if nxtsymb != ')':
                    expression()
                    while nxtsymb == ',':
                        scan()
                        expression()
                    if nxtsymb != ')': error()
                scan()
            elif nxtsymb == '[':
                scan()
                expression()
                while nxtsymb == ',':
                    scan()
                    expression()
                if nxtsymb != ']': error()
                scan()
        elif number() or line(): scan()
        else: error()
        if arithmetic_operation():
            scan()
            expression()

    # число (числовая константа)
    def number():
        return nxtsymb in tokens['N'].values()

    # целое число (числовая константа)
    def integer():
        return nxtsymb in tokens['N'].values()

    # строка (символьная константа)
    def line():
        return nxtsymb in tokens['C'].values()

    # переменная
    def variable():
        if not(name()): error()
        scan()
        if nxtsymb == '[':
            scan()
            expression()
            if nxtsymb != ']': error()
            scan()

    # арифметическая операция
    def arithmetic_operation():
        return nxtsymb in ['%', '*', '+', '-', '/', '+=','-=', '*=', '/=', '%=']

    # составной оператор
    def compound_operator():
        if nxtsymb != '{': error()
        operators()
        if nxtsymb != '}': error()
        scan()

    # оператор присваивания
    def assignment_operator():
        scan()
        variable()
        if nxtsymb != '=': error()
        scan()
        expression()

    # условный оператор
    def conditional_operator():
        if nxtsymb != 'if': error()
        scan()
        if nxtsymb != '(': error()
        condition()
        if nxtsymb != ')': error()
        scan()
        operator()
        if nxtsymb == 'else':
            scan()
            operator()

    # условие
    def condition():
        if unary_log_operation():
            scan()
            if nxtsymb != '(': error()
            log_expression()
            if nxtsymb != ')': error()
            scan()
        else:
            log_expression()
            while binary_log_operation():
                log_expression()

    # унарная логическая операция
    def unary_log_operation():
        return nxtsymb == '!'

    # логическое выражение
    def log_expression():
        scan()
        expression()
        comparison_operation()
        scan()
        expression()

    # операция сравнения
    def comparison_operation():
        return nxtsymb in ['!=', '<', '<=', '==', '>', '>=']

    # бинарная логическая операция
    def binary_log_operation():
        return nxtsymb == '&&' or nxtsymb == '||'

    # цикл for
    def for_loop():
        if nxtsymb != 'for': error()
        scan()
        if nxtsymb != '(': error()
        assignment_operator()
        if nxtsymb != ';': error()
        condition()
        if nxtsymb != ';': error()
        assignment_operator()
        if nxtsymb != ')': error()
        scan()
        operator()

    # цикл while
    def while_loop():
        if nxtsymb != 'while': error()
        scan()
        if nxtsymb != '(': error()
        condition()
        if nxtsymb != ')': error()
        scan()
        operator()

    # оператор break
    def break_operator():
        return nxtsymb == 'break'

    # оператор continue
    def continue_operator():
        return nxtsymb == 'continue'

    # оператор return
    def return_operator():
        if nxtsymb != 'return': error()
        scan()
        expression()

    program()


def process():
    lab1()
    syntax_analyzer()
    dir_path = os.getcwd()
    errorPresent = False
    ername = 'error.txt'
    for file in os.listdir(dir_path+"/script/"):
        if file == ername:
            errorPresent = True
            with open("script/"+ername) as f:
                for line in f:
                    match = re.search(r'\d+', line)
                    if match:
                        line_num = int(match.group())
                        res=open("design/result.txt","w")
                        res.write(str(line_num))
                        res.close()

    filename = 'script.py'
    for file in os.listdir(dir_path+"/script/"):
        if file != filename:
            os.remove(os.path.join(dir_path+"/script/", file))

    if not errorPresent:
        res=open("design/result.txt","w")
        res.write("-1")
        res.close()
        return
    return

process()


