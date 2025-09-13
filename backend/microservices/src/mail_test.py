import mailConection 
from email.mime.text import MIMEText

mailServer = mailConection.InitializeMail()

try:
    sender = 'projectdevicemobile@gmail.com'
    recipient = 'evandromachado2006@gmail.com'

    subject = "E-mail de Teste com Acentuação"
    body = 'Olá, este é um teste de e-mail.'
    msg = MIMEText(body, 'plain', 'utf-8')
    msg['Subject'] = subject
    msg['From'] = sender
    msg['To'] = recipient

    if not mailServer:
        raise Exception("SMTP server is not found")
    mailServer.sendmail(sender, recipient, msg.as_string())
    print("E-mail sent successfully!")
except Exception as e:
    print(f"Fail to sending email: {e}")
