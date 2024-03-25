class ArrayList
{
private:
  int *A;
  int size;
  int length;

public:
  ArrayList();  // default constructor
  ~ArrayList(); // destructor
  void display();
  void append(int value);
};