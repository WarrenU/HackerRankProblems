def is_matched(expression):
    paren_stack = []
    paren_mappings = {
        '[': ']',
        '{': '}',
        '(': ')'
    }
    for c in expression:
        if c in paren_mappings:
            paren_stack.append(paren_mappings[c])
        elif not paren_stack or paren_stack.pop() != c: 
            return False
    return not paren_stack

t = int(input().strip())
for a0 in range(t):
    expression = input().strip()
    if is_matched(expression) == True:
        print("YES")
    else:
        print("NO")
