# Program-that-translates-from-English-to-Pig-Latin


Правила

Вопреки своему названию, поросячья латынь никак не связана с настоящей. 
«Перевод» с английского языка осуществляется следующим образом:
    
Если слово начинается на один или несколько согласных звуков, первые согласные перемещаются в конец слова и добавляется ay. Так ball («шар», «мяч») превращается в all-bay, button («пуговица», «кнопка») — в utton-bay, star («звезда») — в ar-stay, three («три») — в ee-thray, question («вопрос») — в uestion-qay.
    
Если слово начинается на гласный звук, в конце просто добавляется "ay".

## EXAMPLES:  

* They lived happily in the forest afterwards. => Eythay ivedlay appilyhay inay ethay orestfay afterwardsay.
* word => ordway
* expression => expressionay
* Word => Ordway
* Expression => Expressionay
* :word, => :ordway,

## USAGE:

* `git clone git://github.com/0x008800/Program-that-translates-from-English-to-Pig-Latin ~/go/src/Program-that-translates-from-English-to-Pig-Latin`
* `cd ~/go/src/Program-that-translates-from-English-to-Pig-Latin`
* `go build`
* `./english_to_pig_latin path_to_your_filename`

or just run: `go run english_to_pig_latin.go path_to_your_filename`

After running it will produce the file named: "filename(pig-latin)"

also you can...
##RUN TEST: 

* `go test`