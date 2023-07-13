# Dano Project 
----------
- [x] Utils: 
  - [x] getName :
    ```go
      func getName(in int, prefix string) string {
    	_len := len(nameArr)
    	if in <= 0 {
    		return ""
    	}
    	if in > 0 && in <= _len {
    		return fmt.Sprintf("%s%c", prefix, nameArr[in-1])
    	}
    	_min, _max, _radix := _calculateBoundaries(in)
    	var _str string = fmt.Sprint(prefix)
    	var _sub int
    	if _min <= in && in <= _max {
    		_sub = (in - _radix)
    		var i int = 1
    		for i = 1; _sub > _min-1; i++ {
    			_sub = _sub - (_min - 1)
    		}
    		_pref := fmt.Sprintf("%s%c", prefix, nameArr[i-1])
    		return getName(_sub, _pref)
    	}
    	return _str
    }
    ```
  - [ ] giom GUI:

- [x] Define Batches 
  - [x] Define Batch:

    A Batch is a set of Nodes 

    Each node is a struct contains 
    - Name (Name of the node)
    - AR (Average Readings)
    - IR (Instantebous Readings)
    - N (Number of neighbors)
    
    - [x] Define Node