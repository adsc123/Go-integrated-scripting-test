run arr.str{
    str name = Arg(`--name`)
    int times = Arg(`-t`, 0)
    arr.str outputs
    outputs += str(times)
	// str name = ReadString(`Enter your name: `)
	str output = `Hello, %{ ?(*name>0, name, "world") }!`
	outputs += output
	return outputs
}

// The backtick ` is used when evaluations are conducted within a string. Single quotes = char Double = String 
// % = expression 
// ? = conditional operator ?(logical condition, true value, false value)