# Pull base image
FROM python:3.6

# Set environment variables
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

# Set work directory
WORKDIR /thuocsi

# Install dependencies
COPY requirements.txt /thuocsi/
RUN pip install -r requirements.txt

# Copy project
COPY . /thuocsi/
