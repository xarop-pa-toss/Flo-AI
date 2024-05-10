# OpieFlo-AI
A voiced personal assistant split that uses Google's DialogFlow and OpenAIs ChatGPT. It will use multiple RESTful APIs to operate many different services, for example Google Calendar.

The app starts by taking in voice commands in .mp3 format and transcribing them to text using an STT service (Vosk). DialogFlow then decides what to do with the command and utilizes ChatGPT to make conversation!

## Flo
Flo has a female voice and is linked directly to DialogFlow API which tracks Intent, Context, etc and can be used as a memory system for future conversations, thus creating an understanding of the users request habits and becoming capable of activating specific functions based on the user's requests.

Opie is Flo's male counterpart and is much simpler. It basically takes voice, converts to text, sends it to ChatGPT through the OpenAI API and speaks the reply out! One cool thing about ChatGPT service is that you can give it an initial prompt to give it some preset personality and preferences. It will however have access to DialogFlow's memory which will be used as context for it's replies.

An example:
Asking for "Create a dentist's appointment for the 5th of May" would make DialogFlow find the "Create" and "Appointment" keywords which would specify Intent. From there it finds CreateAppointmentInGoogleCalendar which matches the given Intent. It then takes "dentist" and "5th of May" and uses those as contextual arguments for that function and writes to the calendar using the Google Calendar API and voice-replies with whatever information (either success or some error).
DialogFlow is also capable of recognizing that it doesn't have all the information it needs and can immediately ask for clarification. In this example this would happen if the user said "Create a dentists appointment"; Flo would ask for the date.


