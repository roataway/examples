# Descriere
Aici găsești exemple *simple*, care ilustrează cum să te conectezi la server și să primești
datele telemetrice de la transportul public din Chișinău, prin sistemul Roataway.

Obiectivul principal este de a oferi un start rapid, o temelie, deasupra căreia poți
construi ceva propriu.

Fiecare mapă conține un exemplu minimalist, scris într-un anumit limbaj de programare. Tot
acolo vei găsi un readme cu instrucțiuni specifice, dacă e necesar.

# Lista exemplelor

- `javascript-websocket` - o pagină HTML cu JavaScript, care se conectează la server prin
MQTT over websockets, și afișează mesajele telemetrice.
- `javascript-websocket-leaflet` - o versiune mai elaborată a exemplului precedent, care
vizualizează vehiculele pe hartă.
- `python` - conexiune prin MQTT over TCP, cu afișarea pe ecran a mesajelor recepționate.



# Cum să contribui
Dacă vrei să incluzi în lista exemplelor un nou limbaj:

1. Adaugă o mapă cu denumirea limbajului și păstrează acolo fișierele necesare
2. Asigură-te că exemplul funcționează corect
3. Scrie un readme cu instrucțiuni care arată ce trebuie să facă un om, ca să lanseze exemplul
tău, *pornind de la zero*
4. Asigură-te că exemplul e minimalist, elimină totul ce nu e necesar; ține cont de faptul că
publicul țintă e alcătuit din începători
5. Comentează codul, astfel încât să fie clar ce se petrece în locurile-cheie a acestuia
6. Nu recurge la ”clever hacks” pe care nu le va înțelege un începător
7. Nu folosi ”über-enterprise patterns” care prind bine când codul e scris pentru companii cu
sute de angajați, dar care prind nu_bine când codul îl citește un puști care abia învată
programarea
8. Include în readme o secțiune cu referințe la materiale adiționale (screencasts, articole,
cărți, etc.), ca începătorii să-și poată dezvolta deprinderi bune

Când e gata, deschide un pull-request (PR).