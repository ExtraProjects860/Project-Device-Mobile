import smtplib
from src.services.mailConection import InitializeMail
from email.mime.text import MIMEText

mailServer: smtplib.SMTP = InitializeMail("smtp.gmail.com", 587)

try:
    sender: str = 'projectdevicemobile@gmail.com'
    recipient: str = 'evandromachado2006@gmail.com'

    subject: str = "E-mail de Teste com Acentuação"
    body: str = 'Olá, este é um teste de e-mail.'
    msg: MIMEText = MIMEText(body, 'plain', 'utf-8')
    msg['Subject'] = subject
    msg['From'] = sender
    msg['To'] = recipient

    if not mailServer:
        raise Exception("SMTP server is not found")
    mailServer.sendmail(sender, recipient, msg.as_string())
    print("E-mail sent successfully!")
except Exception as e:
    print(f"Fail to sending email: {e}")
