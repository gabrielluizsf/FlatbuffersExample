import express  from 'express'
import { Person }  from './flatbuffers-example/person.js'
import flatbuffers  from 'flatbuffers';

const app = express();
const port = 3000;

import cors from 'cors';  

app.use(cors({ 
    origin: 'http://localhost:3001',
    methods: 'GET'
}));

app.get('/getPerson',(req, res) => {
    try {
        const builder = new flatbuffers.Builder(0);

        const name = builder.createString('Alice');

        Person.startPerson(builder);
        Person.addName(builder, name);
        Person.addAge(builder, 25);
        const personOffset = Person.endPerson(builder);

        builder.finish(personOffset);

        const buffer = builder.asUint8Array();
        const person = Person.getRootAsPerson(new flatbuffers.ByteBuffer(buffer));

        const serializedPerson = {
            name: person.name(),
            age: person.age()
        };

        res.json(serializedPerson);  
    } catch (error) {
        console.error(error);
        res.status(500).json({ Error: error.message });
    }
});

app.listen(port, () => {
    console.log(`Server listening on port ${port}`);
});
