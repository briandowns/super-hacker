package javascript

// http://www.java2s.com/Tutorials/Javascript/Javascript_Data_Structure/0340__Javascript_List.htm

const List = `
function List() { /* w  w  w  .java 2s . c o m*/
    this.listSize = 0; 
    this.pos = 0; 
    this.listData = []; // initializes an empty array to store list elements 
    this.clear = clear; 
    this.find = find; 
    this.toString = toString; 
    this.insert = insert; 
    this.append = append; 
    this.remove = remove; 
    this.front = front; 
    this.end = end; 
    this.prev = prev; 
    this.next = next; 
    this.length = length; 
    this.currentPosition = currentPosition; 
    this.moveTo = moveTo; 
    this.getElement = getElement; 
    this.length = length; 
    this.contains = contains; 
} 
//Adding an Element to a List 
//appends a new element onto the list at the next available position, 
//which will be equal to the value of the listSize variable: 
//After the element is appended, listSize is incremented by 1. 
function append(element) { 
    this.listData[this.listSize++] = element; 
} 

//find() for finding the element to remove: 
//The find function simply iterates through listData looking for the specified element. 
function find(element) { 
    for (var i = 0; i < this.listData.length; ++i) { 
        if (this.listData[i] == element) { 
            return i; 
        } 
    } 
    return -1; 
} 

//Removing an Element from a List 
//we use the splice() mutator function. 
//The remove() function uses the position returned by find() to splice the listData 
//array at that place. 
//After the array is modified, listSize is decremented by 1 to reflect 
//the new size of the list. 
//The function returns true if an element is removed, and false 
//otherwise. Here is the code: 
function remove(element) { 
    var foundAt = this.find(element); 
    if (foundAt > -1) { 
        this.listData.splice(foundAt,1); 
        --this.listSize; 
        return true; 
    } 
    return false; 
} 
//Determining the Number of Elements in a List 
//The length() function returns the number of elements in a list: 
function length() { 
    return this.listSize; 
} 
//Retrieving a List's Elements 
function toString() { 
    return this.listData; 
} 

//Insert: Inserting an Element into a List 
function insert(element, after) { 
    var insertPos = this.find(after); 
    if (insertPos > -1) { 
        this.listData.splice(insertPos+1, 0, element); 
        ++this.listSize; 
        return true; 
    } 
    return false; 
} 
//Clear: Removing All Elements from a List 
function clear() { 
    delete this.listData; 
    this.listData = []; 
    this.listSize = this.pos = 0; 
} 
//Contains: Determining if a Given Value Is in a List 
function contains(element) { 
    for (var i = 0; i < this.listData.length; ++i) { 
        if (this.listData[i] == element) { 
            return true; 
        } 
    } 
    return false; 
} 
//Traversing a List 
function front() { 
    this.pos = 0; 
} 
function end() { 
    this.pos = this.listSize-1; 
} 
function prev() { 
    if (this.pos > 0) { 
        --this.pos; 
    } 
} 
function next() { 
    if (this.pos < this.listSize-1) { 
        ++this.pos; 
    } 
} 
function currentPosition() { 
    return this.pos; 
} 
function moveTo(position) { 
    this.pos = position; 
} 
function getElement() { 
    return this.listData[this.pos]; 
} 

var names = new List(); 
names.append("C"); 
names.append("R"); 
names.append("B"); 
console.log(names.toString()); 
names.remove("R"); 
console.log(names.toString()); 
names.append("C"); 
names.append("R"); 
names.append("C"); 
names.append("J"); 
names.append("B"); 
names.append("D"); 

//move to the first element of the list and display it: 
names.front(); 
console.log(names.getElement());

//move forward one element and display the element's value: 
names.next(); 
console.log(names.getElement());

//move forward twice and backward once, displaying the current element to 
//demonstrate how the prev() function works: 
names.next(); 
names.next(); 
names.prev(); 
console.log(names.getElement());
`
